package formatter

import "time"

type Markdown struct {
	Activities []Activity
}

type Activity struct {
	Event   string
	Message *string
	Repo    string
	Date    time.Time
}
