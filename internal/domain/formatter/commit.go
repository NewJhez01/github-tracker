package formatter

import "time"

type Commit struct {
	Message string
	Name    string
	Email   string
	Date    time.Time
}
