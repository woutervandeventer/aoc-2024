package day6

import (
	"testing"

	"github.com/woutervandeventer/aoc-2024"
)

func TestDistinctGuardPositions(t *testing.T) {
	aoc.TestSolution(t, aoc.OpenFile(t, "example.txt"), DistinctGuardPositions, 41)
}

func TestObstructionPositions(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		aoc.TestSolution(t, aoc.OpenFile(t, "example.txt"), ObstructionPositions, 6)
	})

	t.Run("the real answer is definitely not 2249 or 2248", func(t *testing.T) {
		if got := ObstructionPositions(aoc.OpenFile(t, "input.txt")); got == 2249 || got == 2248 {
			t.Fail()
		}
	})
}
