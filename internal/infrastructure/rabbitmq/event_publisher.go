package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"

	"NewJhez01/github-tracker/internal/domain/formatter"

	"github.com/rabbitmq/amqp091-go"
)

type WorkQueue struct {
	ch    *amqp091.Channel
	queue *amqp091.Queue
}

func NewPublisher(conn *amqp091.Connection) *WorkQueue {
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("failed to connect")
	}

	q := createQueue(ch)
	return &WorkQueue{
		ch:    ch,
		queue: &q,
	}
}

func (r *WorkQueue) Publish(qb *formatter.QueueBody, ctx context.Context) {
	fmt.Println("sent message")
	body, err := json.Marshal(qb)
	err = r.ch.PublishWithContext(ctx,
		"",           // exchange
		r.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		fmt.Println("failed to pub queue")
	}
}

func (r *WorkQueue) Consume() <-chan amqp091.Delivery {
	msg, err := r.ch.Consume(
		r.queue.Name,
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
