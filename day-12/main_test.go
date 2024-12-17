package main

import (
	"os"
	"testing"
)

func TestStones(t *testing.T) {
	f, err := os.Open("example.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { f.Close() })
	want := 1930
	if got := totalPrice(f); got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
