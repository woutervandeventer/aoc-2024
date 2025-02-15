package day12_test

import (
	"testing"

	"github.com/woutervandeventer/aoc-2024"
	day12 "github.com/woutervandeventer/aoc-2024/day-12"
)

func TestTotalPrice(t *testing.T) {
	if got, want := day12.TotalPrice(aoc.OpenFile(t, "example.txt")), 1930; got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
