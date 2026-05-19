package command

import (
	"context"

	"NewJhez01/github-tracker/internal/domain"
)

func SendReport(b []byte, cr domain.CacheRepo, e domain.Smtp) {
	ctx := context.Background()
	d := string(b)
	r := cr.Get(ctx, d)
	e.Send("Github activities for: "+d, r)
}
