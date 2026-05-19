package rabbitmq

import (
	"context"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

type EventPublisher struct {
	Ch    *amqp091.Channel
	Queue *amqp091.Queue
}

func NewPublisher(conn *amqp091.Connection) *EventPublisher {
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("failed to connect")
	}

	q := createQueue(ch)
	return &EventPublisher{
		Ch:    ch,
		Queue: &q,
	}
}

func (r EventPublisher) Publish(s string, ctx context.Context) {
	fmt.Println("sent message")
	err := r.Ch.PublishWithContext(ctx,
		"",           // exchange
		r.Queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(s),
		})
	if err != nil {
		fmt.Println("failed to pub queue")
	}
}
