package day1

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"
)

func TotalDistance(left, right []int) (distance int) {
	a := make([]int, len(left))
	b := make([]int, len(right))
	copy(a, left)
	copy(b, right)
	slices.Sort(a)
	slices.Sort(b)
	for i := range a {
		distance += diff(a[i], b[i])
	}
	return distance
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func SimilarityScore(left, right []int) (score int) {
	rightcount := make(map[int]int)
	for _, n := range right {
		rightcount[n]++
	}
	for _, n := range left {
		score += n * rightcount[n]
	}
	return score
}

func ReadLists(r io.Reader) (a, b []int) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		astr, bstr, _ := strings.Cut(scanner.Text(), "   ")
		aint, _ := strconv.Atoi(astr)
		bint, _ := strconv.Atoi(bstr)
		a = append(a, aint)
		b = append(b, bint)
	}
	return a, b
}
