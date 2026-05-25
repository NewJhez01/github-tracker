package rabbitmq

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"NewJhez01/github-tracker/internal/domain/formatter"

	"github.com/rabbitmq/amqp091-go"
)

type WorkQueue struct {
	ch    *amqp091.Channel
	queue *amqp091.Queue
}

func NewPublisher(conn *amqp091.Connection) (*WorkQueue, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, errors.New("failed to connect to rabbitmq" + err.Error())
	}

	q, err := createQueue(ch)
	if err != nil {
		return nil, errors.New("failed to create work queue prev: " + err.Error())
	}
	return &WorkQueue{
		ch:    ch,
		queue: q,
	}, nil
}

func (r *WorkQueue) Publish(qb *formatter.QueueBody, ctx context.Context) error {
	body, err := json.Marshal(qb)
	if err != nil {
		return errors.New("failed to parse json prev: " + err.Error())
	}
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
		return errors.New("failed to send message prev: " + err.Error())
	}
	log.Println("sent message")
	return nil
}

func (r *WorkQueue) Consume() (<-chan amqp091.Delivery, error) {
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
		return nil, errors.New("failed to consume message prev: " + err.Error())
	}

	log.Println("received message")
	return msg, nil
}
