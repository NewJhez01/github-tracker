package infrastructure

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ParseFileByLine(f *os.File) chan string {
	ch := make(chan string)
	b := make([]byte, 8)
	chunk := ""
	go func() {
		defer close(ch)
		for {
			n, err := f.Read(b)
			if n != 0 {
				data := append([]byte(chunk), b[:n]...)
				lineSplitter(data, ch)
			}
			if err == io.EOF {
				if chunk != "" {
					ch <- chunk
				}
				f.Close()
				return
			}
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}()
	return ch
}

func lineSplitter(b []byte, ch chan string) string {
	rest := ""
	if !strings.Contains(string(b), "\n") {
		return rest
	}
	endsWithNewLine := strings.HasSuffix(string(b), "\n")
	parts := strings.Split(string(b), "\n")
	iterator := parts

	if !endsWithNewLine {
		lastidx := len(parts) - 1
		iterator = parts[:lastidx]
		rest = parts[lastidx]
	}

	for _, v := range iterator {
		ch <- v
	}
	return rest
}
