package query

import (
	"fmt"
	"io"
	"os"

	"NewJhez01/github-tracker/src/internal/services"
)

func FetchFile(repos chan string) {
	f, err := os.Open("conf/repos.toml")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	b := make([]byte, 8)
	chunk := ""
	for {
		n, err := f.Read(b)
		if err != nil && err != io.EOF {
			fmt.Println(err.Error())
		}
		if n != 0 {
			data := append([]byte(chunk), b[:n]...)
			chunk = services.ParseRepos(data, repos)
		}
		if err == io.EOF {
			if chunk != "" {
				repos <- string(chunk)
			}
			close(repos)
			break
		}
	}
}
