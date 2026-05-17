package message

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func send() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("failed to connect to rabbit mq")
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("failed to receive channel")
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		true,    // durability
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		amqp091.Table{
			amqp091.QueueTypeArg: amqp091.QueueTypeQuorum,
		},
	)
	if err != nil {
		fmt.Println("failed to read queue")
	}

	fmt.Println(q)
}
