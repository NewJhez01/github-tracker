package formatter

import (
	"fmt"
	"strings"
	"time"
)

func CreateReport(c []Commit, r string) string {
	repo := strings.Split(r, "/")[1]
	var b strings.Builder

	for _, v := range c {
		fmt.Fprintf(&b, "GITHUB ACTIVITIES IN REPO: %s AT: %s\n\n", repo, time.Now().Format("2006-01-02"))

		fmt.Fprintln(&b, "=====================")
		fmt.Fprintf(&b, "author:   %s\n", v.Name)
		fmt.Fprintf(&b, "message: %s\n", v.Message)
		fmt.Fprintf(&b, "email: %s\n", v.Email)
		fmt.Fprintf(&b, "date:    %s\n", v.Date.Format("2006-01-02 15:04"))
		fmt.Fprintln(&b)
	}

	return b.String()
}
