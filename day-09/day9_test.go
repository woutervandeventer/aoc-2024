package day9_test

import (
	"testing"

	"github.com/woutervandeventer/aoc-2024"
	day9 "github.com/woutervandeventer/aoc-2024/day-09"
)

func TestChecksum(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		if got, want := day9.Checksum(aoc.OpenFile(t, "example.txt")), 2858; got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("real input", func(t *testing.T) {
		if got, want := day9.Checksum(aoc.OpenFile(t, "input.txt")), 6349492251099; got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
