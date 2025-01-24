package day19_test

import (
	"testing"

	"github.com/woutervandeventer/aoc-2024"
	day19 "github.com/woutervandeventer/aoc-2024/day-19"
)

func TestPossibleTowelCombinations(t *testing.T) {
	if got, want := day19.PossibleTowelCombinations(aoc.OpenFile(t, "example.txt")), 16; got != want {
		t.Errorf("got %d possible designs, want %d", got, want)
	}
	if got, want := day19.PossibleTowelCombinations(aoc.OpenFile(t, "input.txt")), 723524534506343; got != want {
		t.Errorf("got %d possible designs, want %d", got, want)
	}
}
