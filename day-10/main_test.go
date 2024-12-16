package main

import (
	"os"
	"testing"
)

func TestScores(t *testing.T) {
	f, err := os.Open("example.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { f.Close() })
	want := 36
	if got := scores(f); got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestRatings(t *testing.T) {
	f, err := os.Open("example.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { f.Close() })
	want := 81
	if got := ratings(f); got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
