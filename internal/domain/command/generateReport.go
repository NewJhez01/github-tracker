package command

import (
	"fmt"
	"strings"

	"NewJhez01/github-tracker/internal/domain/formatter"
)

func GenreateReport(b []byte, s string) {
	// 1 parse the request body
	lines := strings.Split(string(b), "\n")
	// 2 create the markdown
	m := formatter.Markdown{}
	m.CreateReportLines(lines)
	r := m.CreateReport(s)
	fmt.Println(r)
	// 3 call repository func to cache the markdown
	// 4 dispatch the message for the message handler to async handle it
}
