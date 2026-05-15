package formatter

import (
	"fmt"
	"strings"
	"time"
)

func (m *Markdown) CreateReport() string {
	var b strings.Builder

	fmt.Fprintf(&b, "GITHUB ACTIVITIES: %s\n\n", time.Now().Format("2006-01-02"))

	for _, v := range m.Activities {
		fmt.Fprintln(&b, "=====================")
		fmt.Fprintf(&b, "repo:    %s\n", v.Repo)
		fmt.Fprintf(&b, "event:   %s\n", v.Event)
		if v.Message != nil {
			fmt.Fprintf(&b, "message: %s\n", *v.Message)
		}
		fmt.Fprintf(&b, "date:    %s\n", v.Date.Format("2006-01-02 15:04"))
		fmt.Fprintln(&b)
	}

	return b.String()
}
