package formatter

import (
	"strings"
	"testing"
	"time"
)

func TestCreateReport(t *testing.T) {
	fixedTime := time.Date(2026, 5, 20, 19, 20, 0, 0, time.Local)
	commits := []Commit{
		{Message: "hello world", Name: "foo", Email: "foo@bar.com", Date: fixedTime},
		{Message: "goodbye world", Name: "bar", Email: "bar@foo.com", Date: fixedTime},
	}

	result := CreateReport(commits)

	if !strings.Contains(result, "author:   foo") {
		t.Fatalf("missing author foo")
	}
	if !strings.Contains(result, "date:    2026-05-20 19:20") {
		t.Fatalf("wrong date format")
	}
	if !strings.Contains(result, "=====================") {
		t.Fatalf("missing separator")
	}
	if strings.Count(result, "=====================") != 2 {
		t.Fatalf("expected 2 commits, got %d", strings.Count(result, "====================="))
	}
}
