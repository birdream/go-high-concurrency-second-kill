package main

import (
	"fmt"
	"rbmq/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	rbq := RabbitMQ.NewRabbitSimple("imooc-sub-pub")

	for i := 0; i < 100; i++ {
		rbq.PublishPub("subscribe the " + strconv.Itoa(i) + " message")
		fmt.Printf("%d st message published\n", i)
		time.Sleep(time.Second)
	}
}
