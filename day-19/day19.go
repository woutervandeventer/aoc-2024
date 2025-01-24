package day19

import (
	"bufio"
	"io"
	"strings"
	"unicode/utf8"
)

func PossibleDesignCount(input io.Reader) (count int) {
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	towels := strings.Split(scanner.Text(), ", ")

	// Discard empty line
	scanner.Scan()

	for scanner.Scan() {
		design := scanner.Text()

		var designPossible func(design string) bool
		designPossible = func(design string) bool {
			if len(design) == 0 {
				return true
			}
			for _, towel := range towels {
				if strings.HasPrefix(design, towel) {
					if designPossible(design[utf8.RuneCountInString(towel):]) {
						return true
					}
				}
			}
			return false
		}

		if designPossible(design) {
			count++
		}
	}

	return count
}
