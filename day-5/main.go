package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	correct, incorrect := sumAllMiddlePageNos(os.Stdin)
	fmt.Printf("correct: %d, incorrect: %d\n", correct, incorrect)
}

func sumAllMiddlePageNos(r io.Reader) (correct, incorrect int) {
	// Read the rules
	rules := make(map[string]struct{ before, after []string })
	scanner := bufio.NewScanner(r)
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
		leftrule.before = append(leftrule.before, right)
		rightrule.after = append(rightrule.after, left)
		rules[left], rules[right] = leftrule, rightrule
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Read the updates
	var updates [][]string
	for scanner.Scan() {
		updates = append(updates, strings.Split(scanner.Text(), ","))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var correctUpdates, incorrectUpdates [][]string
Updates:
	for _, update := range updates {
		for i, page := range update {
			for j := i + 1; j < len(update); j++ {
				otherPage := update[j]
				rulesForPage, rulesForOtherPage := rules[page], rules[otherPage]
				if !slices.Contains(rulesForPage.before, otherPage) || !slices.Contains(rulesForOtherPage.after, page) {
					incorrectUpdates = append(incorrectUpdates, update)
					continue Updates
				}
			}
		}
		correctUpdates = append(correctUpdates, update)
	}

	// Sort incorrect updates using the rules
	for _, update := range incorrectUpdates {
		slices.SortFunc(update, func(a, b string) int {
			if slices.Contains(rules[a].before, b) || slices.Contains(rules[b].after, a) {
				return -1
			}
			return 1
		})
	}

	return sumMiddlePageNos(correctUpdates), sumMiddlePageNos(incorrectUpdates)
}

func sumMiddlePageNos(updates [][]string) (count int) {
	for _, update := range updates {
		i, err := strconv.Atoi(update[len(update)/2])
		if err != nil {
			panic(err)
		}
		count += i
	}
	return count
}
