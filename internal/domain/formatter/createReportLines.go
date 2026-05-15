package formatter

import "time"

func (m *Markdown) CreateReportLines(l []string) {
	for _, v := range l {
		m.Activities = append(m.Activities, Activity{
			Event:   v,
			Message: &v,
			Repo:    v,
			Date:    time.Now(),
		},
		)
	}
}
