package main

import (
	"os"
	"testing"
)

func TestUniqueAntinodes(t *testing.T) {
	f, err := os.Open("example.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { f.Close() })
	want := 34
	if got := uniqueAntinodes(f); got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
