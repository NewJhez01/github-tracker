package command

import (
	"context"
	"fmt"
	"time"

	"NewJhez01/github-tracker/internal/domain/formatter"
	"NewJhez01/github-tracker/internal/infrastructure/parser"
	"NewJhez01/github-tracker/internal/infrastructure/rabbitmq"
	"NewJhez01/github-tracker/internal/repo"
)

func GenerateReport(b []byte, s string, since time.Time) {
	if string(b) == "[]" {
		fmt.Println("No data for day")
	}
	c, err := parser.ParseJson(b)
	if err != nil {
		fmt.Println("parser func failed")
	}
	r := formatter.CreateReport(c, s)
	ctx := context.Background()
	yesterday := since.Format("2006-01-02")
	repo.Set(ctx, r, yesterday)
	rabbitmq.Send(yesterday)
}
