package day3

import (
	"io"
	"strconv"
)

type parser struct {
	scanner scanner
}

func newParser(r io.Reader) parser {
	return parser{
		scanner: newScanner(r),
	}
}

func (p parser) parse() (pg program) {
	enabled := true
	for {
		tok, _ := p.scanner.scan()
		switch tok {
		case eof:
			return pg
		case do:
			if p.parseParens() {
				enabled = true
			}
		case dont:
			if p.parseParens() {
				enabled = false
			}
		case mul:
			if !enabled {
				continue
			}
			m, ok := p.parseMul()
			if !ok {
				continue
			}
			pg.muls = append(pg.muls, m)
		}
	}
}

func (p parser) parseParens() bool {
	if tok, _ := p.scanner.scan(); tok != lparen {
		return false
	}
	tok, _ := p.scanner.scan()
	return tok == rparen
}

func (p parser) parseMul() (multiplication, bool) {
	if tok, _ := p.scanner.scan(); tok != lparen {
		return multiplication{}, false
	}
	tok, astr := p.scanner.scan()
	if tok != number {
		return multiplication{}, false
	}
	if tok, _ := p.scanner.scan(); tok != comma {
		return multiplication{}, false
	}
	tok, bstr := p.scanner.scan()
	if tok != number {
		return multiplication{}, false
	}
	if tok, _ = p.scanner.scan(); tok != rparen {
		return multiplication{}, false
	}
	a, err := strconv.Atoi(astr)
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(bstr)
	if err != nil {
		panic(err)
	}
	return multiplication{a: a, b: b}, true
}
