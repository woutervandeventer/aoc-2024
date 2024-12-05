package main

import (
	"os"
	"testing"
)

func TestSumAllMiddlePageNos(t *testing.T) {
	f, err := os.Open("example.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { f.Close() })
	_ = f
	correct, incorrect := sumAllMiddlePageNos(f)
	if correct != 143 {
		t.Errorf("got %d want 143", correct)
	}
	if incorrect != 123 {
		t.Errorf("got %d want 123", incorrect)
	}
}
