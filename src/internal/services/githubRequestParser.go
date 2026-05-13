package services

import (
	"fmt"
	"strings"
)

func ParseRequest(b []byte) {
	s := strings.Split(string(b), ",")
	for _, v := range s {
		if strings.Contains(v, "message") || strings.Contains(v, "url") {
			fmt.Println(v)
		}
	}
}
