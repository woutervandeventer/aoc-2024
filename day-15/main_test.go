package main

import (
	"testing"

	"github.com/woutervandeventer/aoc-2024"
)

func TestSolution(t *testing.T) {
	aoc.TestSolution(t, aoc.OpenFile(t, "example.txt"), solution, 10092)
}
