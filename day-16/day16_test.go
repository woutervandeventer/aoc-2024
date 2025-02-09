package day16_test

import (
	"testing"

	"github.com/woutervandeventer/aoc-2024"
	day16 "github.com/woutervandeventer/aoc-2024/day-16"
)

func TestLowestScore(t *testing.T) {
	t.Run("first example", func(t *testing.T) {
		if got, want := day16.LowestScore(aoc.OpenFile(t, "example.txt")), 7036; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("second example", func(t *testing.T) {
		if got, want := day16.LowestScore(aoc.OpenFile(t, "example2.txt")), 11048; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
