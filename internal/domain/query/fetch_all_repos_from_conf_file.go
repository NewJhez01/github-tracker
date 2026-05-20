package query

import (
	"errors"
	"os"

	"NewJhez01/github-tracker/internal/domain"
)

const REPO_FILE_PATH = "conf/repos.toml"

func FetchRepos(fParser domain.FileParser) (chan string, error) {
	f, err := os.Open(REPO_FILE_PATH)
	if err != nil {
		return nil, errors.New("failed to open file" + err.Error())
	}
	return fParser.ParseFileByLine(f), nil
}
