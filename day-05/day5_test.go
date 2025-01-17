package day5

import (
	"testing"

	"github.com/woutervandeventer/aoc-2024"
)

func TestSumAllMiddlePageNos(t *testing.T) {
	correct, incorrect := SumMiddlePageNos(aoc.OpenFile(t, "example.txt"))
	if correct != 143 {
		t.Errorf("got %d want 143", correct)
	}
	if incorrect != 123 {
		t.Errorf("got %d want 123", incorrect)
	}
}
