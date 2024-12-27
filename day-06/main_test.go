package main

import (
	"testing"

	"github.com/woutervandeventer/aoc-2024"
)

func TestDistinctPositions(t *testing.T) {
	aoc.TestSolution(t, aoc.OpenFile(t, "example.txt"), distinctPositions, 41)
}

func TestObstructionPositions(t *testing.T) {
	aoc.TestSolution(t, aoc.OpenFile(t, "example.txt"), obstructionPositions, 6)
}
