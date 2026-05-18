package command

import (
	"context"
	"fmt"

	"NewJhez01/github-tracker/internal/repo"
)

func SendReport(b []byte) {
	ctx := context.Background()
	fmt.Println(repo.Get(ctx, string(b)))
}
