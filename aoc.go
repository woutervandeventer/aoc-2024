package aoc

import (
	"io"
	"os"
	"testing"
)

func TestSolution(t *testing.T, input io.Reader, solution func(io.Reader) int, answer int) {
	t.Helper()
	if got := solution(input); got != answer {
		t.Errorf("got %d, want %d", got, answer)
	}
}

func OpenFile(t *testing.T, name string) *os.File {
	t.Helper()
	f, err := os.Open(name)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { f.Close() })
	return f
}
