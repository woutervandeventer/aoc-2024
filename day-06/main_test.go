package day6

import (
	"testing"

	"github.com/woutervandeventer/aoc-2024"
)

func TestDistinctGuardPositions(t *testing.T) {
	aoc.TestSolution(t, aoc.OpenFile(t, "example.txt"), DistinctGuardPositions, 41)
}

func TestObstructionPositions(t *testing.T) {
	aoc.TestSolution(t, aoc.OpenFile(t, "example.txt"), ObstructionPositions, 6)
}
