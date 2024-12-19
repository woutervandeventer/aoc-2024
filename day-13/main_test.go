package main

import (
	"os"
	"testing"
)

func TestFewestTokens(t *testing.T) {
	f, err := os.Open("example.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { f.Close() })
	want := 480
	if got := fewestTokens(f); got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
