package main

import (
	"NewJhez01/github-tracker/internal/handler/http"
	"NewJhez01/github-tracker/internal/handler/message"
)

func main() {
	go message.Send()
	http.FetchGithubData()
	select {}
}
