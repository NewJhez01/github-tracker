package parser

import (
	"fmt"
	"testing"
	"time"
)

func TestParseJson(t *testing.T) {
	date := time.Now().Format(time.RFC3339)
	set1 := fmt.Sprintf(`{"commit": {
		"message": "hello world",
		"committer": {
			"name": "foo",
			"email": "foo@gmail.com",
			"date": "%s"
		}
	}}`, date)
	set2 := fmt.Sprintf(`{"commit": {
		"message": "goodbye world",
		"committer": {
			"name": "bar",
			"email": "bar@gmail.com",
			"date": "%s"
		}
	}}`, date)

	testData := fmt.Sprintf("[%s, %s]", set1, set2)

	gp := NewGithubParser()
	commits, err := gp.ParseJson([]byte(testData))
	if err != nil {
		t.Fatalf("function produced an unexpected error")
	}
	if len(commits) != 2 {
		t.Fatalf("function provided invalid result")
	}

	result1 := commits[0]
	if result1.Name != "foo" || result1.Email != "foo@gmail.com" || result1.Message != "hello world" || result1.Date.Format(time.RFC3339) != date {
		t.Fatalf("function provided invalid struct")
	}

	result2 := commits[1]
	if result2.Name != "bar" || result2.Email != "bar@gmail.com" || result2.Message != "goodbye world" || result2.Date.Format(time.RFC3339) != date {
		t.Fatalf("function provided invalid struct")
	}
}

func TestParseJsonError(t *testing.T) {
	falseJson := "{hello world}"
	gp := NewGithubParser()
	_, err := gp.ParseJson([]byte(falseJson))
	if err == nil {
		t.Fatalf("expected an err")
	}
}
