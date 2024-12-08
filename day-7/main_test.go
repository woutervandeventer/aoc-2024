package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestTotalCalibrationResult(t *testing.T) {
	f, err := os.Open("example.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { f.Close() })
	b, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	want := 3749
	if got := totalCalibrationResult(bytes.NewReader(b), add, mul); got != want {
		t.Errorf("got %d want %d", got, want)
	}

	withConcat := 11387
	if got := totalCalibrationResult(bytes.NewReader(b), add, mul, concat); got != withConcat {
		t.Errorf("with concat: got %d want %d", got, withConcat)
	}
}
