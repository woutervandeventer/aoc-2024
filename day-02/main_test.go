package main

import (
	"testing"

	"github.com/woutervandeventer/aoc-2024"
)

func TestSafeReportsWithDampener(t *testing.T) {
	aoc.TestSolution(t, aoc.OpenFile(t, "example.txt"), safeReportsWithDampener, 4)
}
