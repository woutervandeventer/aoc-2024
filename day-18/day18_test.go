package day18_test

import (
	"testing"

	"github.com/woutervandeventer/aoc-2024"
	day18 "github.com/woutervandeventer/aoc-2024/day-18"
)

func TestMinimumStepsToExit(t *testing.T) {
	if got, want := day18.MinimumStepsToExit(6, 12, aoc.OpenFile(t, "example.txt")), 22; got != want {
		t.Errorf("got %d minimum steps, want %d", got, want)
	}
}

func TestBlockingByteCoordinates(t *testing.T) {
	if got, want := day18.BlockingByteCoordinates(6, 12, aoc.OpenFile(t, "example.txt")), "6,1"; got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
