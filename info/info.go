package info

import "encoding/json"

type ServiceInfo struct {
	Prefix string `json:"prefix"`
	Key    string `json:"key"`
	Name   string `json:"name"`
	Addr   string `json:"addr"`
}

func New(prefix, key, name, addr string) *ServiceInfo {
	return &ServiceInfo{
		Prefix: prefix,
		Key:    key,
		Name:   name,
		Addr:   addr,
	}
}

func (s *ServiceInfo) String() string {
	tmp, _ := json.Marshal(s)
	return string(tmp)
}

func Marshall(val []byte) *ServiceInfo {
	s := new(ServiceInfo)
	json.Unmarshal(val, &s)
	return s
}
