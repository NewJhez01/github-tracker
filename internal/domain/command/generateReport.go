package command

import (
	"fmt"
	"strings"

	"NewJhez01/github-tracker/internal/domain/formatter"
)

func GenreateReport(b []byte) {
	// 1 parse the request body
	lines := strings.Split(string(b), "\n")
	// 2 create the markdown
	m := formatter.Markdown{}
	m.CreateReportLines(lines)
	s := m.CreateReport()
	fmt.Println(s)
	// 3 call repository func to cache the markdown
	// 4 dispatch the message for the message handler to async handle it
}
