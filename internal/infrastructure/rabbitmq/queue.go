package rabbitmq

import (
	"errors"

	"github.com/rabbitmq/amqp091-go"
)

func createQueue(ch *amqp091.Channel) (*amqp091.Queue, error) {
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
	if err != nil {
		return nil, errors.New("failed to connect to rabbitmq: " + err.Error())
	}

	return &q, nil
}
