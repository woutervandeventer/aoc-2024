package day5

import (
	"bufio"
	"bytes"
	"io"
	"slices"
	"strconv"
	"strings"
)

func SumMiddlePageNos(r io.Reader) (correct, incorrect int) {
	rules := readRules(r)
	// Why can't you share a reader between multiple scanners?
	updateScanner := newUpdateScanner(r)

	for updateScanner.scan() {
		update := updateScanner.update()
		switch isCorrect(update, rules) {
		case true:
			correct += middlePageNo(update)
		case false:
			sort(update, rules)
			incorrect += middlePageNo(update)
		}
	}

	return correct, incorrect
}

type page = string

type rules map[page]struct{ isBefore, isAfter []page }

func readRules(r io.Reader) rules {
	scanner := bufio.NewScanner(r)
	rules := make(rules)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { // Reached the empty line in the middle
			break
		}
		left, right, ok := strings.Cut(line, "|")
		if !ok {
			panic("no \"|\" in line: " + line)
		}
		leftrule, rightrule := rules[left], rules[right]
		leftrule.isBefore = append(leftrule.isBefore, right)
		rightrule.isAfter = append(rightrule.isAfter, left)
		rules[left], rules[right] = leftrule, rightrule
	}
	return rules
}

type updateScanner struct {
	scanner *bufio.Scanner
	buf     []page
}

func newUpdateScanner(r io.Reader) updateScanner {
	return updateScanner{
		scanner: bufio.NewScanner(r),
	}
}

func (s updateScanner) scan() bool {
	return s.scanner.Scan()
}

func (s updateScanner) update() []page {
	clear(s.buf)
	for _, b := range bytes.Split(s.scanner.Bytes(), []byte(",")) {
		s.buf = append(s.buf, page(b))
	}
	return s.buf
}

func isCorrect(update []page, rules rules) bool {
	for i, page := range update {
		for j := i + 1; j < len(update); j++ {
			otherPage := update[j]
			rulesForPage, rulesForOtherPage := rules[page], rules[otherPage]
			if !slices.Contains(rulesForPage.isBefore, otherPage) || !slices.Contains(rulesForOtherPage.isAfter, page) {
				return false
			}
		}
	}
	return true
}

func sort(update []page, rules rules) {
	slices.SortFunc(update, func(a, b string) int {
		if slices.Contains(rules[a].isBefore, b) || slices.Contains(rules[b].isAfter, a) {
			return -1
		}
		return 1
	})
}

func middlePageNo(update []page) (n int) {
	n, err := strconv.Atoi(update[len(update)/2])
	if err != nil {
		panic(err)
	}
	return n
}
