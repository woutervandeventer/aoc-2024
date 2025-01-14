package day2

import (
	"testing"

	"github.com/woutervandeventer/aoc-2024"
)

func TestSafeReportsWithDampener(t *testing.T) {
	aoc.TestSolution(t, aoc.OpenFile(t, "example.txt"), SafeReportsWithDampener, 4)
	aoc.TestSolution(t, aoc.OpenFile(t, "input.txt"), SafeReportsWithDampener, 700)
}
