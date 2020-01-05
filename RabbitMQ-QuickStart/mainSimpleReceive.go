package main

import "RabbitMQ"

func main() {
	rbmq := RabbitMQ.NewRabbitMQSimple("imoocSimple")
	rbmq.ConsumeSimple()
}
