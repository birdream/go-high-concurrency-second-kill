package main

import "rbmq/RabbitMQ"

func main() {
	imoocTwo := RabbitMQ.NewRabbitMQRouting("exImooc", "imooc_two")
	imoocTwo.ReceiveRouting()
}
