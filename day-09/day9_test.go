package day9_test

import (
	"os"
	"testing"

	day9 "github.com/woutervandeventer/aoc-2024/day-09"
)

func TestChecksum(t *testing.T) {
	f, err := os.Open("example.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { f.Close() })
	want := 1928
	if got := day9.Checksum(f); got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
