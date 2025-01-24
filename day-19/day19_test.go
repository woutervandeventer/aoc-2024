package day19_test

import (
	"testing"

	"github.com/woutervandeventer/aoc-2024"
	day19 "github.com/woutervandeventer/aoc-2024/day-19"
)

func TestPossibleDesignCount(t *testing.T) {
	if got, want := day19.PossibleDesignCount(aoc.OpenFile(t, "example.txt")), 6; got != want {
		t.Errorf("got %d possible designs, want %d", got, want)
	}
}
