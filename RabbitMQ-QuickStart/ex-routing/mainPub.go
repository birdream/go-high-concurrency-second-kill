package main

import (
	"fmt"
	"rbmq/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	imoocOne := RabbitMQ.NewRabbitMQRouting("exImooc", "imooc_one")
	imoocTwo := RabbitMQ.NewRabbitMQRouting("exImooc", "imooc_two")

	for i := 0; i < 100; i++ {
		imoocOne.PublishRouting("Hello one!" + strconv.Itoa(i))
		imoocTwo.PublishRouting("Hello two!" + strconv.Itoa(i))

		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
