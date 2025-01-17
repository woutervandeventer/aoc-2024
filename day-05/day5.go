package day5

import (
	"bufio"
	"bytes"
	"io"
	"slices"
	"strconv"
	"strings"
)

// SumMiddlePageNos returns the sum of the middle page numbers for both the
// correct and incorrect updates.
func SumMiddlePageNos(r io.Reader) (correct, incorrect int) {
	scanner := bufio.NewScanner(r)
	rules := readRules(scanner)

	for scanner.Scan() {
		update := parseUpdate(scanner.Bytes())
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

type rules map[page]struct{ before, after []page }

func readRules(scanner *bufio.Scanner) rules {
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
		leftrule.after = append(leftrule.after, right)
		rightrule.before = append(rightrule.before, left)
		rules[left], rules[right] = leftrule, rightrule
	}
	return rules
}

func parseUpdate(b []byte) (update []page) {
	for _, p := range bytes.Split(b, []byte(",")) {
		update = append(update, page(p))
	}
	return update
}

func isCorrect(update []page, rules rules) bool {
	for i, page := range update {
		for j := i + 1; j < len(update); j++ {
			otherPage := update[j]
			rulesForPage, rulesForOtherPage := rules[page], rules[otherPage]
			if !slices.Contains(rulesForPage.after, otherPage) || !slices.Contains(rulesForOtherPage.before, page) {
				return false
			}
		}
	}
	return true
}

func sort(update []page, rules rules) {
	slices.SortFunc(update, func(a, b string) int {
		if slices.Contains(rules[a].after, b) || slices.Contains(rules[b].before, a) {
			return -1
		}
		return 1
	})
}

func middlePageNo(update []page) int {
	n, err := strconv.Atoi(update[len(update)/2])
	if err != nil {
		panic(err)
	}
	return n
}
