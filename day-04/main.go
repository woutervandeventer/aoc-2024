package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
	"strings"
)

var (
	xmasrexp = regexp.MustCompile("XMAS")
	masrexp  = regexp.MustCompile("MAS")
)

func main() {
	var xmas, xMas bool
	flag.BoolVar(&xmas, "xmas", false, "Find all XMAS occurences")
	flag.BoolVar(&xMas, "x-mas", false, "Find all X-MAS occurences")
	flag.Parse()
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	switch {
	case xmas:
		fmt.Println(countxmas(string(input)))
	case xMas:
		fmt.Println(countxMas(string(input)))
	default:
		fmt.Println("Choose a mode")
	}
}

func countxmas(input string) (count int) {
	var lines [][]byte // All possible lines within the input

	var matrix [][]byte
	for _, line := range strings.Split(input, "\n") {
		if line != "" {
			matrix = append(matrix, []byte(line))
		}
	}

	// Append the rows
	lines = append(lines, matrix...)

	// Append the columns
	for x := range matrix[0] {
		var col []byte
		for _, row := range matrix {
			col = append(col, row[x])
		}
		lines = append(lines, col)
	}

	// Append the diagonals
	lines = append(lines, diagonals(matrix)...)
	slices.Reverse(matrix) // Reversing the matrix yields the perpendicular diagonals
	lines = append(lines, diagonals(matrix)...)

	for _, line := range lines {
		count += countMatches(line, xmasrexp)
	}

	return count
}

func diagonals(matrix [][]byte) (diags [][]byte) {
	// Upper left
	for i := range matrix[0] {
		var diag []byte
		for x, y := i, 0; x >= 0; x, y = x-1, y+1 {
			diag = append(diag, matrix[y][x])
		}
		diags = append(diags, diag)
	}
	// Lower right exluding middle
	for row := 1; row < len(matrix); row++ {
		var diag []byte
		for x, y := len(matrix[0])-1, row; y < len(matrix); x, y = x-1, y+1 {
			diag = append(diag, matrix[y][x])
		}
		diags = append(diags, diag)
	}
	return diags
}

func countxMas(input string) (count int) {
	var matrix [][]byte
	for _, line := range strings.Split(input, "\n") {
		if line != "" {
			matrix = append(matrix, []byte(line))
		}
	}
	for y, row := range matrix {
		for x, char := range row {
			if char != 'A' {
				continue
			}
			// Check if char is not at an edge of the matrix.
			if x == 0 || x == len(matrix[y])-1 || y == 0 || y == len(matrix)-1 {
				continue
			}
			diag1 := []byte{matrix[y-1][x-1], matrix[y][x], matrix[y+1][x+1]}
			if !match(diag1, masrexp) {
				continue
			}
			diag2 := []byte{matrix[y+1][x-1], matrix[y][x], matrix[y-1][x+1]}
			if !match(diag2, masrexp) {
				continue
			}
			count++
		}
	}
	return count
}

func countMatches(in []byte, rexp *regexp.Regexp) (count int) {
	count += len(rexp.FindAll(in, -1))
	slices.Reverse(in)
	count += len(rexp.FindAll(in, -1))
	return count
}

func match(in []byte, rexp *regexp.Regexp) bool {
	if rexp.Match(in) {
		return true
	}
	slices.Reverse(in)
	return rexp.Match(in)
}
