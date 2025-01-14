package day3

import "testing"

func TestAddMuls(t *testing.T) {
	cases := []struct {
		input  string
		answer int
	}{
		{"mul(2,2)", 4},
		{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", 161},
	}
	for _, c := range cases {
		if got := AddMuls(c.input); got != c.answer {
			t.Errorf("got %d want %d", got, c.answer)
		}
	}
}

func TestEnablesAddMuls(t *testing.T) {
	cases := []struct {
		input  string
		answer int
	}{
		{"mul(2,2)", 4},
		{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", 48},
	}
	for _, c := range cases {
		if got := AddEnabledMuls([]byte(c.input)); got != c.answer {
			t.Errorf("got %d want %d", got, c.answer)
		}
	}
}
