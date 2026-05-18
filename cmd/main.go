package main

import (
	"fmt"

	"NewJhez01/github-tracker/internal/handler/http"
	"NewJhez01/github-tracker/internal/handler/message"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load env")
	}
	go message.Send()
	http.FetchGithubData()
	select {}
}
