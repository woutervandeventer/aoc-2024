package main

import (
	"os"
	"testing"
)

func TestCountCorrectMiddlePageNos(t *testing.T) {
	f, err := os.Open("example.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { f.Close() })
	_ = f
	if got := countCorrectMiddlePageNos(f); got != 143 {
		t.Errorf("got %d want 143", got)
	}
}
