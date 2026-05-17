package command

import (
	"fmt"

	"NewJhez01/github-tracker/internal/domain/formatter"
	"NewJhez01/github-tracker/internal/infrastructure"
	"NewJhez01/github-tracker/internal/repo"
)

func GenerateReport(b []byte, s string) {
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
	repo.Get()
	// 4 dispatch the message for the message handler to async handle it
}
