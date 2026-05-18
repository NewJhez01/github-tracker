package message

import (
	"fmt"

	"NewJhez01/github-tracker/internal/domain/command"

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
		"send_email", // name
		true,         // durability
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		amqp091.Table{
			amqp091.QueueTypeArg: amqp091.QueueTypeQuorum,
		},
	)

	msg, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		fmt.Println("failed to fetch messages")
	}
	command.SendReport(msg)
}
