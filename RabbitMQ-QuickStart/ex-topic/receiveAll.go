package main

import "rbmq/RabbitMQ"

func main() {
	imoocOne := RabbitMQ.NewRabbitMQTopic("exImoocTopic", "#")
	imoocOne.ReceiveTopic()
}
