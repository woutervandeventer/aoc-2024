package main

import (
	"strings"
	"testing"
)

func TestCountxmas(t *testing.T) {
	cases := []struct {
		input string
		count int
	}{
		{
			input: `
______X
_____M_
____A__
___S___
__A____
_M_____
X______`,
			count: 2,
		},
		{
			input: `
X______
_M_____
__A____
___S___
____A__
_____M_
______X`,
			count: 2,
		},
		{
			input: `
X_____X
_M___M_
__A_A__
___S___
__A_A__
_M___M_
X_____X`,
			count: 4,
		},
		{
			input: `
X__X__X
_M_M_M_
__AAA__
XMASAMX
__AAA__
_M_M_M_
X__X__X`,
			count: 8,
		},
		{
			input: `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			count: 18,
		},
	}
	for _, tc := range cases {
		wonl, _ := strings.CutPrefix(tc.input, "\n")
		if got := countxmas(wonl); got != tc.count {
			t.Errorf("got %d want %d", got, tc.count)
		}
	}
}

func TestCountxMas(t *testing.T) {
	cases := []struct {
		input string
		count int
	}{
		{
			input: `
M.S
.A.
M.S`,
			count: 1,
		},
		{
			input: `
.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`,
			count: 9,
		},
	}
	for _, tc := range cases {
		wonl, _ := strings.CutPrefix(tc.input, "\n")
		if got := countxMas(wonl); got != tc.count {
			t.Errorf("got %d want %d", got, tc.count)
		}
	}
}
