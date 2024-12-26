package main

import (
	"os"
	"testing"
)

func TestSafetyFactor(t *testing.T) {
	f, err := os.Open("example.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { f.Close() })
	want := 12
	if got := safetyFactor(f, 11, 7); got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
