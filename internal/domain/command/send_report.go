package command

import (
	"context"
	"encoding/json"
	"fmt"

	"NewJhez01/github-tracker/internal/domain"
	"NewJhez01/github-tracker/internal/domain/formatter"
)

func SendReport(b []byte, cr domain.CacheRepo, e domain.Smtp) {
	qb := formatter.QueueBody{}
	err := json.Unmarshal(b, &qb)
	if err != nil {
		fmt.Println("failed to serialize json from message")
	}
	ctx := context.Background()
	r := cr.Get(ctx, qb.Date)
	subject := "Github activities for: " + qb.Date + " in repo: " + qb.Repo
	e.Send(subject, r)
}
