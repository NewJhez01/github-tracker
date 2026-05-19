package message

import (
	"NewJhez01/github-tracker/internal/domain"
	"NewJhez01/github-tracker/internal/domain/command"
)

func Send(r domain.RabbitMq, cr domain.CacheRepo, smtp domain.Smtp) {
	msg := r.Consume()
	for m := range msg {
		command.SendReport(m.Body, cr, smtp)
	}
}
