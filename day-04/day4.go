package day4

import (
	"bufio"
	"io"
	"strings"
)

func CountXmas(input io.Reader) (count int) {
	const xmas = "XMAS"

	matrix := readWordSearch(input)
	width, height := len(matrix[0]), len(matrix)

	// Rows
	for _, row := range matrix {
		count += matches(string(row), xmas)
	}

	// Columns
	for x := 0; x < width; x++ {
		col := make([]rune, height)
		for y := 0; y < height; y++ {
			col[y] = matrix[y][x]
		}
		count += matches(string(col), xmas)
	}

	// Diagonals bottom left to upper right
	for y := 0; y < height; y++ {
		var diag []rune
		for x, y := 0, y; x < width && y >= 0; x, y = x+1, y-1 {
			diag = append(diag, matrix[y][x])
		}
		count += matches(string(diag), xmas)
	}
	for x := 1; x < width; x++ {
		var diag []rune
		for x, y := x, height-1; x < width && y >= 0; x, y = x+1, y-1 {
			diag = append(diag, matrix[y][x])
		}
		count += matches(string(diag), xmas)
	}

	// Diagonals upper left to bottom right
	for y := height - 1; y >= 0; y-- {
		var diag []rune
		for x, y := 0, y; x < width && y < height; x, y = x+1, y+1 {
			diag = append(diag, matrix[y][x])
		}
		count += matches(string(diag), xmas)
	}
	for x := 1; x < width; x++ {
		var diag []rune
		for x, y := x, 0; x < width && y < height; x, y = x+1, y+1 {
			diag = append(diag, matrix[y][x])
		}
		count += matches(string(diag), xmas)
	}

	return count
}

func CountXMas(input io.Reader) (count int) {
	const mas = "MAS"

	matrix := readWordSearch(input)

	// Skip the edges, because the X cannot occur there.
	for y := 1; y < len(matrix)-1; y++ {
		for x := 1; x < len(matrix[0])-1; x++ {
			ch := matrix[y][x]
			if ch != 'A' {
				continue
			}
			diag1 := []rune{matrix[y-1][x-1], ch, matrix[y+1][x+1]}
			if matches(string(diag1), mas) == 0 {
				continue
			}
			diag2 := []rune{matrix[y+1][x-1], ch, matrix[y-1][x+1]}
			if matches(string(diag2), mas) == 0 {
				continue
			}
			count++
		}
	}

	return count
}

func readWordSearch(input io.Reader) [][]rune {
	var matrix [][]rune
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		matrix = append(matrix, []rune(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return matrix
}

func matches(s, word string) (n int) {
	reversed := reverse(word)
	for i := 0; i <= len(s)-len(word); i++ {
		if strings.HasPrefix(s[i:], word) || strings.HasPrefix(s[i:], reversed) {
			n++
		}
	}
	return n
}

func reverse(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		j := len(runes) - 1 - i
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
