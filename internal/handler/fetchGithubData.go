package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"NewJhez01/github-tracker/internal/domain/query"
	"NewJhez01/github-tracker/internal/services"
)

func main() {
	ch := make(chan string)
	go query.FetchFile(ch)
	for v := range ch {
		fmt.Println(v)
		req, err := http.NewRequest("GET", "https://api.github.com/repos/"+v+"/commits?since=2026-05-07", nil)
		if err != nil {
			fmt.Println("fail")
		}
		req.Header.Set("Accept", "application/vnd.github+json")
		req.Header.Set("X-GitHub-Api-Version", "2026-03-10")

		c := &http.Client{Timeout: time.Duration(1) * time.Second}

		resp, err := c.Do(req)
		if err != nil {
			fmt.Println("failed to get response")
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		services.ParseRequest(body)
	}
}
