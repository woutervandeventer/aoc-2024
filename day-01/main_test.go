package main

import (
	"testing"

	"github.com/woutervandeventer/aoc-2024"
)

func TestTotalDistance(t *testing.T) {
	f := aoc.OpenFile(t, "example.txt")
	got := totalDistance(readLists(f))
	want := 11
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSimilarityScore(t *testing.T) {
	f := aoc.OpenFile(t, "example.txt")
	got := similarityScore(readLists(f))
	want := 31
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
