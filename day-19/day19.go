package day19

import (
	"bufio"
	"io"
	"strings"
	"unicode/utf8"
)

func PossibleTowelCombinations(input io.Reader) (count int) {
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	towels := strings.Split(scanner.Text(), ", ")

	// Discard empty line
	scanner.Scan()

	// Dynamic programming: store solved subproblems in a map
	combinations := make(map[string]int)

	for scanner.Scan() {
		design := scanner.Text()

		var possibleCombinations func(design string) int
		possibleCombinations = func(design string) (n int) {
			if combs, exists := combinations[design]; exists {
				return combs
			}
			if len(design) == 0 {
				return 1
			}
			for _, towel := range towels {
				if !strings.HasPrefix(design, towel) {
					continue
				}
				rest := design[utf8.RuneCountInString(towel):]
				combs := possibleCombinations(rest)
				combinations[rest] = combs
				n += combs
			}
			return n
		}

		count += possibleCombinations(design)
	}

	return count
}
