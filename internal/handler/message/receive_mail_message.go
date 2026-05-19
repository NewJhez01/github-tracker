package message

import (
	"fmt"

	"NewJhez01/github-tracker/internal/domain"
	"NewJhez01/github-tracker/internal/domain/command"
	"NewJhez01/github-tracker/internal/infrastructure/rabbitmq"

	"github.com/rabbitmq/amqp091-go"
)

func Send(r rabbitmq.EventPublisher, cr domain.CacheRepo, smtp domain.Smtp) {
	msg, err := r.Ch.Consume(
		r.Queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		amqp091.Table{
			amqp091.QueueTypeArg: amqp091.QueueTypeQuorum,
		},
	)
	if err != nil {
		fmt.Println("failed to fetch messages")
	}
	for m := range msg {
		command.SendReport(m.Body, cr, smtp)
	}
}
