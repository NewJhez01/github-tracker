package query

import (
	"fmt"
	"os"

	"NewJhez01/github-tracker/internal/domain"
)

const REPO_FILE_PATH = "conf/repos.toml"

func FetchRepos(fParser domain.FileParser) chan string {
	f, err := os.Open(REPO_FILE_PATH)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return fParser.ParseFileByLine(f)
}
