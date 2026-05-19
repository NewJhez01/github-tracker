package command

import (
	"context"
	"fmt"
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
) {
	if string(b) == "[]" {
		fmt.Println("No data for day")
	}
	c, err := p.ParseJson(b)
	if err != nil {
		fmt.Println("parser func failed")
	}
	r := formatter.CreateReport(c, s)
	ctx := context.Background()
	yesterday := since.Format("2006-01-02")
	cr.Set(ctx, r, yesterday)
	rabbitMQ.Publish(yesterday, ctx)
}
