package command

import (
	"context"

	"NewJhez01/github-tracker/internal/infrastructure/email"
	"NewJhez01/github-tracker/internal/repo"
)

func SendReport(b []byte) {
	ctx := context.Background()
	r := (repo.Get(ctx, string(b)))
	// to do handle secrets
	c := email.CreateSmtpClient("**************")
	c.Send("Github activities for today", r)
}
