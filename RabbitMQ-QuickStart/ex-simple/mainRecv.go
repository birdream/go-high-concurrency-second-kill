package main

import "rbmq/RabbitMQ"

func main() {
	rbmq := RabbitMQ.NewRabbitSimple("imoocSimple")
	rbmq.ConsumeSimple()
}
