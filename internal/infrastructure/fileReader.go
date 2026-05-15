package infrastructure

import (
	"errors"
	"os"
)

func ReadFile(p string) (*os.File, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, errors.New("Invalid File path")
	}

	return f, nil
}
