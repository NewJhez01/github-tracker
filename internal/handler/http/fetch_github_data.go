package http

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"NewJhez01/github-tracker/internal/domain"
	"NewJhez01/github-tracker/internal/domain/command"
	"NewJhez01/github-tracker/internal/domain/query"
)

func FetchGithubData(
	p domain.JsonParser,
	rabbitMq domain.RabbitMq,
	fParser domain.FileParser,
	cr domain.CacheRepo,
) {
	ch := query.FetchRepos(fParser)
	since := time.Now().Add(-48 * time.Hour)
	c := &http.Client{Timeout: time.Duration(1) * time.Second}
	for v := range ch {
		req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/repos/%s/commits?since=%s", v, since.Format("2006-01-02")), nil)
		if err != nil {
			fmt.Println("fail")
		}
		req.Header.Set("Accept", "application/vnd.github+json")
		req.Header.Set("X-GitHub-Api-Version", "2026-03-10")

		resp, err := c.Do(req)
		if err != nil {
			fmt.Println("failed to get response")
			continue
		}

		if resp.StatusCode != 200 {
			fmt.Println("statuscode: ", resp.StatusCode)
			continue
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		command.GenerateReport(p, rabbitMq, body, v, since, cr, v)
	}
}
