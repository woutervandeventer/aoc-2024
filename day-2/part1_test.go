package main

import (
	"fmt"
	"testing"
)

func TestIsSafe(t *testing.T) {
	cases := []struct {
		report []int
		safe   bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, true},
		{[]int{1, 3, 6, 7, 9}, true},
		{[]int{82, 84, 85, 87, 90, 92, 91, 91}, false},
		{[]int{77, 80, 81, 84, 85, 86, 86, 90}, false},
		{[]int{71, 69, 70, 71, 72, 75}, true},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v should be %t", c.report, c.safe), func(t *testing.T) {
			if isSafe(c.report) != c.safe {
				t.Fail()
			}
		})
	}
}
