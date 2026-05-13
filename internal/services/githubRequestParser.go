package services

import (
	"fmt"
	"strings"
)

func ParseRequest(b []byte) {
	s := strings.Split(string(b), ",")
	for _, v := range s {
		if (strings.Contains(v, "message") || strings.Contains(v, "repos")) && (!strings.Contains(v, "repos_url") && !strings.Contains(v, "comments_url")) {
			fmt.Println(v)
		}
	}
}
