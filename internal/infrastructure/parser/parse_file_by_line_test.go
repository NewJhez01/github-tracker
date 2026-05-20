package parser

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestParseFileByLine(t *testing.T) {
	fileFixture := "foo\nbar\nxar"
	expected := []string{"foo", "bar", "xar"}
	fPath := filepath.Join(os.TempDir(), "test")
	f, err := os.Create(fPath)
	_, err = f.Write([]byte(fileFixture))
	if err != nil {
		log.Fatalf("failed to create temp file for test")
	}
	fp := NewFileParser()
	result := fp.ParseFileByLine(f)
	i := 0
	for v := range result {
		if expected[i] != v {
			t.Fatalf("value does not match expected")
		}
		i++
	}
}

type readerCloserError struct{}

func (readerCloserError) Read(_ []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func (readerCloserError) Close() error {
	return nil
}

func TestParseFileByLineError(t *testing.T) {
	fp := NewFileParser()
	result := fp.ParseFileByLine(readerCloserError{})
	i := 0
	for range result {
		i++
	}

	if i != 0 {
		t.Fatalf("Programm should have exited on reader error")
	}
}
