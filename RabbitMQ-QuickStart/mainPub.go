package main

import (
	"fmt"
	"rbmq/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitSimple("" +
		"imoocSimple")
	rabbitmq.PublishSimple("Hello imooc!")
	fmt.Println("发送成功！")
}
