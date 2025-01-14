package day3

import (
	"io"
)

func AddMuls(input io.Reader) int {
	parser := newParser(input)
	program := parser.parse()
	return program.calculateMultiplications()
}
