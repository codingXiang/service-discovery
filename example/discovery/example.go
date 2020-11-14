package main

import (
	"github.com/codingXiang/go-logger/v2"
	"github.com/codingXiang/service-discovery/discovery"
	"github.com/codingXiang/service-discovery/util"
	"log"
	"time"
)

func main() {
	var endpoints = []string{"localhost:2379"}
	ser := discovery.New(&util.ETCDAuth{
		Endpoints: endpoints,
		Username:  "root",
		Password:  "a12345",
	})
	defer ser.Close()
	logger.Log.Info(ser.GetServiceValue("/backend/service/example"))
	ser.WatchService("/backend/service")
	for {
		select {
		case <-time.Tick(10 * time.Second):
			log.Println(ser.GetServices())
		}
	}
}