package register

import (
	"context"
	"github.com/codingXiang/go-logger/v2"
	"github.com/codingXiang/service-discovery/info"
	"github.com/coreos/etcd/clientv3"
	"time"
)

//ServiceRegister 创建租约注册服务
type ServiceRegister struct {
	cli           *clientv3.Client //etcd client
	leaseID       clientv3.LeaseID //租约ID
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	key           string            //key
	val           *info.ServiceInfo //value
}

//New 新增註冊服務
func New(endpoints []string, info *info.ServiceInfo, lease int64) (*ServiceRegister, error) {
	if logger.Log == nil {
		logger.Log = logger.Default()
	}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logger.Log.Fatal(err)
		return nil, err
	}

	ser := &ServiceRegister{
		cli: cli,
		key: info.Prefix + info.Key,
		val: info,
	}

	//申请租约设置时间keepalive
	if err := ser.putKeyWithLease(lease); err != nil {
		return nil, err
	}

	return ser, nil
}

//putKeyWithLease 續約
func (s *ServiceRegister) putKeyWithLease(lease int64) error {
	//设置租约时间
	resp, err := s.cli.Grant(context.Background(), lease)
	if err != nil {
		return err
	}
	//注册服务并绑定租约
	_, err = s.cli.Put(context.Background(), s.key, s.val.String(), clientv3.WithLease(resp.ID))
	if err != nil {
		return err
	}
	//设置续租 定期发送需求请求
	leaseRespChan, err := s.cli.KeepAlive(context.Background(), resp.ID)

	if err != nil {
		return err
	}
	s.leaseID = resp.ID
	logger.Log.Debug(s.leaseID)
	s.keepAliveChan = leaseRespChan
	logger.Log.Debug("Put key: ", s.key, " , value: ", s.val, " success")
	return nil
}

//ListenLeaseRespChan 監聽續約狀態
func (s *ServiceRegister) ListenLeaseRespChan() {
	for leaseKeepResp := range s.keepAliveChan {
		logger.Log.Debug("續約成功", leaseKeepResp)
	}
	logger.Log.Debug("關閉續約")
}

//Close 刪除服務
func (s *ServiceRegister) Close() error {
	//刪除服務
	if _, err := s.cli.Revoke(context.Background(), s.leaseID); err != nil {
		return err
	}
	logger.Log.Debug("刪除服務")
	return s.cli.Close()
}
