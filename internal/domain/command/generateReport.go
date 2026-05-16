package command

import (
	"fmt"

	"NewJhez01/github-tracker/internal/domain/formatter"
	"NewJhez01/github-tracker/internal/infrastructure"
)

func GenreateReport(b []byte, s string) {
	// 1 parse the request body
	c, err := infrastructure.ParseJson(b)
	if err != nil {
		fmt.Println("parser func failed")
	}
	// 2 create the markdown
	r := formatter.CreateReport(c, s)
	fmt.Println(r)
	// 3 call repository func to cache the markdown
	// 4 dispatch the message for the message handler to async handle it
}
