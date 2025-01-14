package day3

type token int

const (
	corrupted token = iota
	eof
	ident
	mul
	lparen
	rparen
	number
	comma
)
