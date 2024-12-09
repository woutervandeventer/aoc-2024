package main

import (
	"os"
	"testing"
)

func TestChecksum(t *testing.T) {
	f, err := os.Open("example.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { f.Close() })
	want := 1928
	if got := checksum(f); got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
