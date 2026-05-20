package command

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"NewJhez01/github-tracker/internal/domain"
	"NewJhez01/github-tracker/internal/domain/formatter"
)

func SendReport(b []byte, cr domain.CacheRepo, e domain.Smtp) error {
	qb := formatter.QueueBody{}
	err := json.Unmarshal(b, &qb)
	if err != nil {
		return errors.New("failed to parse json " + err.Error())
	}
	ctx := context.Background()
	r, err := cr.Get(ctx, qb.Date)
	if err != nil {
		return fmt.Errorf("failed to value from redis with key %s, prev: %s", qb.Date, err.Error())
	}
	subject := "Github activities for: " + qb.Date + " in repo: " + qb.Repo
	e.Send(subject, r)
	return nil
}
