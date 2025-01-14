package day2

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func SafeReportsWithDampener(input io.Reader) (safe int) {
	r := newReportScanner(input)
	for r.scan() {
		if isSafeWithDampener(r.report()) {
			safe++
		}
	}
	return safe
}

func isSafeWithDampener(report []int) bool {
	const (
		_ = iota
		ascending
		descending
	)

	var isSafeSkipIndex func(skip int) bool
	isSafeSkipIndex = func(skip int) bool {
		var reportDirection int
		var i, j int
		for {
			if i == skip {
				i++
			}
			j = i + 1
			if j == skip {
				j++
			}
			if j >= len(report) {
				break
			}
			left, right := report[i], report[j]
			allowedToSkip := skip == -1
			if left == right || diff(left, right) > 3 {
				if !allowedToSkip {
					return false
				}
				return isSafeSkipIndex(i) || isSafeSkipIndex(j)
			}
			var direction int
			if right > left {
				direction = ascending
			} else {
				direction = descending
			}
			if reportDirection == 0 {
				reportDirection = direction
			}
			if direction != reportDirection {
				if !allowedToSkip {
					return false
				}
				// Skip the previous entry too, because that could fix the order.
				return isSafeSkipIndex(i-1) || isSafeSkipIndex(i) || isSafeSkipIndex(j)
			}
			i++
		}
		return true
	}

	return isSafeSkipIndex(-1) // Don't skip any index at first.
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
