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
	repo string,
) {
	if string(b) == "[]" {
		fmt.Println("No data for day")
		return
	}
	c, err := p.ParseJson(b)
	if err != nil {
		fmt.Println("parser func failed")
	}
	r := formatter.CreateReport(c, s)
	ctx := context.Background()
	yesterday := since.Format("2006-01-02")
	body := yesterday + " for repo: " + repo
	cr.Set(ctx, r, yesterday)
	rabbitMQ.Publish(body, ctx)
}
