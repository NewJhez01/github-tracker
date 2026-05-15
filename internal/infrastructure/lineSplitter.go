package infrastructure

import "strings"

func SplitLines(chunk []byte, ch chan string) string {
	if !strings.Contains(string(chunk), "\n") {
		return string(chunk)
	}
	rest := ""
	parts := strings.Split(string(chunk), "\n")
	if !strings.Contains((string(chunk[len(chunk)-1])), "\n") {
		rest = parts[len(parts)-1]
		parts = parts[:len(parts)-1]
	}
	for _, v := range parts {
		if v == "" {
			continue
		}
		ch <- v
	}
	return rest
}
