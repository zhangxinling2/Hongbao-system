package main

import (
	"github.com/tietang/go-eureka-client/eureka"
	"time"
)

func main() {
	conf := eureka.Config{DialTimeout: time.Second * 30}
	client := eureka.NewClientByConfig([]string{"http://127.0.0.1:8761/eureka"}, conf)
	appName := "go-example"
	instance := eureka.NewInstanceInfo("test.com", appName, "127.0.0.2", 8080, 30, false)
	client.RegisterInstance(appName, instance)
	client.Start()
	c := make(chan int)
	<-c
}
