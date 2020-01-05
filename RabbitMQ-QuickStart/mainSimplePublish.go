package main

import (
	"RabbitMQ"
	"fmt"
)

func main() {
	rbmq := RabbitMQ.NewRabbitMQSimple("imoocSimple")
	rbmq.PublishSimple("Hello Norman!")
	fmt.Println("publish succeed")
}
