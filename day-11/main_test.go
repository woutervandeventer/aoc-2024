package main

import (
	"os"
	"testing"
)

func TestStones(t *testing.T) {
	f, err := os.Open("example.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { f.Close() })
	want := 55312
	if got := stones(readStones(f), 25); got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func BenchmarkStones(b *testing.B) {
	f, err := os.Open("example.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.Cleanup(func() { f.Close() })
	ss := readStones(f)
	b.ResetTimer()
	for range b.N {
		stones(ss, 25)
	}
}
