package http

import (
	"fmt"
	"io"
	"log"
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
	ch, err := query.FetchRepos(fParser)
	if err != nil {
		log.Fatalf("failed to fetch repos prev: %s", err.Error())
	}
	since := time.Now().Add(-48 * time.Hour)
	c := &http.Client{Timeout: time.Duration(1) * time.Second}
	for v := range ch {
		req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/repos/%s/commits?since=%s", v, since.Format("2006-01-02")), nil)
		if err != nil {
			log.Printf("request for %s failed with error :%s continuing", v, err.Error())
			continue
		}
		req.Header.Set("Accept", "application/vnd.github+json")
		req.Header.Set("X-GitHub-Api-Version", "2026-03-10")

		resp, err := c.Do(req)
		if err != nil {
			log.Printf("response failure %s continuing", err.Error())
			continue
		}

		if resp.StatusCode != 200 {
			log.Printf("unexpected status code %v for request: %s continuing", resp.StatusCode, resp.Body)
			continue
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Printf("failed to read body prev: %s", err.Error())
			continue
		}
		err = command.GenerateReport(p, rabbitMq, body, v, since, cr, v)
		if err != nil {
			log.Printf("failed to generate report prev: %s", err.Error())
		}
	}
}
