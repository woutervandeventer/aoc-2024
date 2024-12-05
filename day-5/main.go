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
	fmt.Println(countCorrectMiddlePageNos(os.Stdin))
}

func countCorrectMiddlePageNos(r io.Reader) (count int) {
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
	for page, rule := range rules {
		fmt.Printf("page %s should come before: %v and after %v\n", page, rule.before, rule.after)
	}

	// Read the updates
	var updates [][]string
	for scanner.Scan() {
		updates = append(updates, strings.Split(scanner.Text(), ","))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var middlePages []string
currentupdate:
	for _, update := range updates {
		for i, page := range update {
			for j := i + 1; j < len(update); j++ {
				otherPage := update[j]
				rulesForPage, rulesForOtherPage := rules[page], rules[otherPage]
				if !slices.Contains(rulesForPage.before, otherPage) || !slices.Contains(rulesForOtherPage.after, page) {
					fmt.Printf("update %v is not valid\n", update)
					continue currentupdate
				}
			}
		}
		fmt.Printf("update %v is valid\n", update)
		middlePages = append(middlePages, update[len(update)/2])
	}

	for _, page := range middlePages {
		i, err := strconv.Atoi(page)
		if err != nil {
			panic(err)
		}
		count += i
	}

	return count
}
