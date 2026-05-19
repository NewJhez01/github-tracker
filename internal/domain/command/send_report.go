package command

import (
	"context"
	"os"

	"NewJhez01/github-tracker/internal/infrastructure/email"
	"NewJhez01/github-tracker/internal/repo"
)

func SendReport(b []byte) {
	ctx := context.Background()
	d := string(b)
	r := repo.Get(ctx, d)
	c := email.CreateSmtpClient(os.Getenv("SMTP_FROM"), os.Getenv("SMTP_HOST"), os.Getenv("SMTP_PASSWORD"))
	c.Send(os.Getenv("SMTP_ADDR"), "Github activities for: "+d, r)
}
