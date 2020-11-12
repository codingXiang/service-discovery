package main

import (
	"github.com/codingXiang/go-logger/v2"
	"github.com/codingXiang/service-discovery/info"
	"github.com/codingXiang/service-discovery/register"
)

func main() {
	var endpoints = []string{"http://localhost:32770"}
	_, err := register.New(endpoints, info.New("/backend/service/", "example", "範例", "http://127.0.0.1:9999"), 5)
	if err != nil {
		logger.Log.Fatal(err)
	}
	//监听续租相应chan
	//go ser.ListenLeaseRespChan()
	select {
	// case <-time.After(20 * time.Second):
	//     ser.Close()
	}
}
