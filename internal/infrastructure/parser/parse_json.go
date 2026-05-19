package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"NewJhez01/github-tracker/internal/domain/formatter"
)

type commit struct {
	Commit commitDetail `json:"commit"`
}

type commitDetail struct {
	Message   string    `json:"message"`
	Committer committer `json:"committer"`
}

type committer struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Date  time.Time `json:"date"`
}

type GithubParser struct{}

func NewGithubParser() *GithubParser {
	return &GithubParser{}
}

func (GithubParser) ParseJson(b []byte) ([]formatter.Commit, error) {
	c := []commit{}
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
