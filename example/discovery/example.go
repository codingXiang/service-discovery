package main

import (
	"github.com/codingXiang/go-logger/v2"
	"github.com/codingXiang/service-discovery/discovery"
	"log"
	"time"
)

func main() {
	var endpoints = []string{"localhost:32770"}
	ser := discovery.New(endpoints)
	defer ser.Close()
	logger.Log.Info(ser.GetServiceValue("/backend/service/example"))
	//ser.WatchService("/gRPC/")
	for {
		select {
		case <-time.Tick(10 * time.Second):
			log.Println(ser.GetServices())
		}
	}
}