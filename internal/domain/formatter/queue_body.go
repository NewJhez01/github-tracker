package formatter

type QueueBody struct {
	Date string `json:"date"`
	Repo string `json:"repo"`
}

func NewQueueBody(d, r string) *QueueBody {
	return &QueueBody{
		Date: d,
		Repo: r,
	}
}
