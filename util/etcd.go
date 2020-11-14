package util

import (
	"go.etcd.io/etcd/clientv3"
	"time"
)

type ETCDAuth struct {
	Endpoints []string
	Username  string
	Password  string
}

func NewETCDClient(auth *ETCDAuth) (*clientv3.Client, error) {
	conf := clientv3.Config{
		Endpoints:   auth.Endpoints,
		DialTimeout: 5 * time.Second,
	}

	if auth.Username != "" && auth.Password != "" {
		conf.Username = auth.Username
		conf.Password = auth.Password
	}

	return clientv3.New(conf)
}
