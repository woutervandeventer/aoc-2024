package day18_test

import (
	"testing"
	"time"

	"github.com/woutervandeventer/aoc-2024"
	day18 "github.com/woutervandeventer/aoc-2024/day-18"
)

func TestMinimumStepsToExit(t *testing.T) {
	if got, want := day18.MinimumStepsToExit(7, 12, aoc.OpenFile(t, "example.txt")), 22; got != want {
		t.Errorf("got %d minimum steps, want %d", got, want)
	}

	gotc := make(chan int)
	go func() { gotc <- day18.MinimumStepsToExit(71, 1024, aoc.OpenFile(t, "input.txt")) }()
	select {
	case got := <-gotc:
		if tooHigh := 542; got >= tooHigh {
			t.Errorf("got %d, has to be lower than %d", got, tooHigh)
		}
	case <-time.After(2 * time.Second):
		t.Error("timeout")
	}
}
