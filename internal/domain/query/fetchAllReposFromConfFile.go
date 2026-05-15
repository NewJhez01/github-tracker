package query

import (
	"fmt"
	"os"

	"NewJhez01/github-tracker/internal/infrastructure"
)

const REPO_FILE_PATH = "conf/repos.toml"

func FetchFile() chan string {
	f, err := os.Open(REPO_FILE_PATH)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return infrastructure.SplitLines(f)
}
