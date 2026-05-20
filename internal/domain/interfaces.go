package domain

import (
	"context"
	"io"

	"NewJhez01/github-tracker/internal/domain/formatter"

	"github.com/rabbitmq/amqp091-go"
)

type JsonParser interface {
	ParseJson(buffer []byte) ([]formatter.Commit, error)
}

type RabbitMq interface {
	Publish(body *formatter.QueueBody, ctx context.Context)
	Consume() <-chan amqp091.Delivery
}

type FileParser interface {
	ParseFileByLine(file io.ReadCloser) chan string
}

type CacheRepo interface {
	Set(ctx context.Context, report, key string)
	Get(ctx context.Context, key string) string
}

type Smtp interface {
	Send(subject, body string)
}
