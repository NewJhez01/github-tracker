package message

import (
	"log"

	"NewJhez01/github-tracker/internal/domain"
	"NewJhez01/github-tracker/internal/domain/command"
)

func Send(r domain.RabbitMq, cr domain.CacheRepo, smtp domain.Smtp) {
	msg, err := r.Consume()
	if err != nil {
		log.Fatalf("Failed to consume message")
	}

	for m := range msg {
		command.SendReport(m.Body, cr, smtp)
	}
}
