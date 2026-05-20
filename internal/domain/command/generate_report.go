package command

import (
	"context"
	"errors"
	"time"

	"NewJhez01/github-tracker/internal/domain"
	"NewJhez01/github-tracker/internal/domain/formatter"
)

func GenerateReport(
	p domain.JsonParser,
	rabbitMQ domain.RabbitMq,
	b []byte, s string,
	since time.Time,
	cr domain.CacheRepo,
	repo string,
) error {
	if string(b) == "[]" {
		return errors.New("No data for given day")
	}
	c, err := p.ParseJson(b)
	if err != nil {
		return errors.New("failed to parse json prev: " + err.Error())
	}
	r := formatter.CreateReport(c)
	ctx := context.Background()
	yesterday := since.Format("2006-01-02")
	qb := formatter.NewQueueBody(yesterday, repo)
	cr.Set(ctx, r, yesterday)
	rabbitMQ.Publish(qb, ctx)
	return nil
}
