package command

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func SendReport(ch <-chan amqp091.Delivery) {
	go func() {
		for v := range ch {
			fmt.Println("consumed message")
			fmt.Println(v)
		}
	}()
}
