package main

import (
	"os"
	"testing"
)

func TestDistinctPositions(t *testing.T) {
	f, err := os.Open("example.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { f.Close() })
	want := 41
	if got := distinctPositions(f); got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}

func TestObstructionPositions(t *testing.T) {
	f, err := os.Open("example.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { f.Close() })
	want := 6
	if got := obstructionPositions(f); got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}
