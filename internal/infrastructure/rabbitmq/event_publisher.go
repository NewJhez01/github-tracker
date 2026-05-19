package rabbitmq

import (
	"context"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

type WorkQueue struct {
	Ch    *amqp091.Channel
	Queue *amqp091.Queue
}

func NewPublisher(conn *amqp091.Connection) *WorkQueue {
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("failed to connect")
	}

	q := createQueue(ch)
	return &WorkQueue{
		Ch:    ch,
		Queue: &q,
	}
}

func (r *WorkQueue) Publish(s string, ctx context.Context) {
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

func (r *WorkQueue) Consume() <-chan amqp091.Delivery {
	msg, err := r.Ch.Consume(
		r.Queue.Name,
		"",
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		amqp091.Table{
			amqp091.QueueTypeArg: amqp091.QueueTypeQuorum,
		},
	)
	if err != nil {
		fmt.Println("fail to recieve message")
	}

	return msg
}
