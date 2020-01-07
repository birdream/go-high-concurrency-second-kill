package main

import "rbmq/RabbitMQ"

func main() {
	rbq := RabbitMQ.NewRabbitMQPubSub("newProduct")
	rbq.RecieveSub()
}
