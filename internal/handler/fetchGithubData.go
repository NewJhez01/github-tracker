package handler

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"NewJhez01/github-tracker/internal/domain/command"
	"NewJhez01/github-tracker/internal/domain/query"
)

func FetchGithubData() {
	ch := query.FetchRepos()
	yesterday := time.Now().Add(-24 * time.Hour).Format("2006-01-02")
	for v := range ch {
		req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/repos/%s/commits?since=%s", v, yesterday), nil)
		if err != nil {
			fmt.Println("fail")
		}
		req.Header.Set("Accept", "application/vnd.github+json")
		req.Header.Set("X-GitHub-Api-Version", "2026-03-10")

		c := &http.Client{Timeout: time.Duration(1) * time.Second}

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
		command.GenerateReport(body, v)
	}
}
