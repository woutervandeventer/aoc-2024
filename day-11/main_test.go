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
	want := 55312
	if got := stones(f, 25); got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
