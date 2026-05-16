package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"NewJhez01/github-tracker/internal/domain/formatter"
)

type Commit struct {
	Commit CommitDetail `json:"commit"`
}

type CommitDetail struct {
	Message   string    `json:"message"`
	Committer Committer `json:"committer"`
}

type Committer struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Date  time.Time `json:"date"`
}

func ParseJson(b []byte) ([]formatter.Commit, error) {
	c := []Commit{}
	err := json.Unmarshal(b, &c)
	if err != nil {
		fmt.Println("marshall error", err)
		return nil, errors.New("fail to marshall json")
	}
	p := []formatter.Commit{}
	for _, v := range c {
		cm := formatter.Commit{
			Message: v.Commit.Message,
			Name:    v.Commit.Committer.Name,
			Email:   v.Commit.Committer.Email,
			Date:    v.Commit.Committer.Date,
		}

		p = append(p, cm)
	}
	return p, nil
}
