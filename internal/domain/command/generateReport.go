package command

import (
	"context"
	"fmt"
	"time"

	"NewJhez01/github-tracker/internal/domain/formatter"
	"NewJhez01/github-tracker/internal/infrastructure"
	"NewJhez01/github-tracker/internal/repo"
)

func GenerateReport(b []byte, s string, since time.Time) {
	if string(b) == "[]" {
		fmt.Println("No data for day")
	}
	c, err := infrastructure.ParseJson(b)
	if err != nil {
		fmt.Println("parser func failed")
	}
	r := formatter.CreateReport(c, s)
	fmt.Println(r)
	// 3 call repository func to cache the markdown
	ctx := context.Background()
	yesterday := since.Format("2006-01-02")
	repo.Set(ctx, s, yesterday)
	repo.Get(ctx, yesterday)
	// 4 dispatch the message for the message handler to async handle it
}
