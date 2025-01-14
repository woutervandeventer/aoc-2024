package day1

import (
	"testing"

	"github.com/woutervandeventer/aoc-2024"
)

func TestTotalDistance(t *testing.T) {
	f := aoc.OpenFile(t, "example.txt")
	got := TotalDistance(ReadLists(f))
	want := 11
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSimilarityScore(t *testing.T) {
	f := aoc.OpenFile(t, "example.txt")
	got := SimilarityScore(ReadLists(f))
	want := 31
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
