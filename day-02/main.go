package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(safeReportsWithDampener(os.Stdin))
}

func safeReportsWithDampener(input io.Reader) (safe int) {
	r := newReportScanner(input)
	for r.scan() {
		if isSafe(r.report()) {
			safe++
		}
	}
	return safe
}

func isSafe(report []int) bool {
	var isSafeR func(tries int, report []int) bool
	isSafeR = func(tries int, report []int) bool {
		if tries == 0 {
			return false
		}
		tries--
		var asc bool
		for i := 0; i < len(report)-1; i++ {
			left, right := report[i], report[i+1]
			currentlyAsc := right > left
			if i == 0 {
				asc = currentlyAsc
			}
			woleft := append(append([]int{}, report[:i]...), report[i+1:]...)
			woright := append(append([]int{}, report[:i+1]...), report[i+2:]...)
			if left == right || diff(left, right) > 3 {
				return isSafeR(tries, woleft) || isSafeR(tries, woright)
			}
			if asc != currentlyAsc {
				woprev := append(append([]int{}, report[:i-1]...), report[i:]...)
				return isSafeR(tries, woprev) || isSafeR(tries, woleft) || isSafeR(tries, woright)
			}
		}
		return true
	}
	return isSafeR(2, report)
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

type reportScanner struct {
	scanner *bufio.Scanner
	rep     []int
}

func newReportScanner(r io.Reader) *reportScanner {
	return &reportScanner{
		scanner: bufio.NewScanner(r),
	}
}

func (r *reportScanner) scan() bool {
	if !r.scanner.Scan() {
		return false
	}
	strs := strings.Split(r.scanner.Text(), " ")
	if r.rep == nil {
		r.rep = make([]int, len(strs))
	}
	for i, s := range strs {
		r.rep[i], _ = strconv.Atoi(s)
	}
	return true
}

func (r *reportScanner) report() []int {
	return r.rep
}
