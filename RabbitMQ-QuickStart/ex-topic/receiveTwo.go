package main

import "rbmq/RabbitMQ"

func main() {
	imoocTwo := RabbitMQ.NewRabbitMQTopic("exImoocTopic", "imooc.*.two")
	imoocTwo.ReceiveTopic()
}
