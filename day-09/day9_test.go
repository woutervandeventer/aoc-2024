package day9_test

import (
	"testing"

	"github.com/woutervandeventer/aoc-2024"
	day9 "github.com/woutervandeventer/aoc-2024/day-09"
)

func TestChecksum(t *testing.T) {
	if got, want := day9.Checksum(aoc.OpenFile(t, "example.txt")), 2858; got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
