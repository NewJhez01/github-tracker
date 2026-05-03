package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	req, err := http.NewRequest("GET", "https://api.github.com/repos/NewJhez01/http_from_scratch/commits", nil)
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
	fmt.Println(string(body))
}
