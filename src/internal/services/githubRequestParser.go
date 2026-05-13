package services

import (
	"fmt"
	"strings"
)

func ParseRequest(b []byte) {
	s := strings.Split(string(b), ",")
	for _, v := range s {
		fmt.Println(v)
	}
}
