package query

import (
	"fmt"
	"io"

	"NewJhez01/github-tracker/internal/infrastructure"
)

const REPO_FILE_PATH = "conf/repos.toml"

func FetchFile(repos chan string) {
	f, err := infrastructure.ReadFile(REPO_FILE_PATH)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	b := make([]byte, 8)
	chunk := ""
	for {
		n, err := f.Read(b)
		if n != 0 {
			data := append([]byte(chunk), b[:n]...)
			chunk = infrastructure.SplitLines(data, repos)
		}
		if err == io.EOF {
			if chunk != "" {
				repos <- chunk
			}
			close(repos)
			return
		}
		if err != nil {
			fmt.Println(err.Error())
			close(repos)
			return
		}
	}
}
