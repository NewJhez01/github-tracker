package services

import (
	"fmt"
	"strings"
)

func ParseRepos(chunk []byte, repos chan string) string {
	if !strings.Contains(string(chunk), "\n") {
		return string(chunk)
	}
	rest := ""
	parts := strings.Split(string(chunk), "\n")
	fmt.Println("parts: ")
	fmt.Println(parts)
	if !strings.Contains((string(chunk[len(chunk)-1])), "\n") {
		rest = parts[len(parts)-1]
		parts = parts[:len(parts)-1]
	}
	for _, v := range parts {
		if v == "" {
			continue
		}
		repos <- v
	}
	return rest
}
