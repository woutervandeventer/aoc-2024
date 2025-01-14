package day3

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type scanner struct {
	reader *bufio.Reader
	buf    *bytes.Buffer
}

func newScanner(r io.Reader) scanner {
	return scanner{
		reader: bufio.NewReader(r),
		buf:    new(bytes.Buffer),
	}
}

func (s scanner) scan() (token, string) {
	ch := s.read()
	if ch == 0 {
		return eof, ""
	}
	if isLetter(ch) {
		s.unread()
		return s.readIdent()
	}
	if isNumber(ch) {
		s.unread()
		return s.readNumber()
	}
	switch ch {
	case 0:
		return eof, ""
	case '(':
		return lparen, string(ch)
	case ')':
		return rparen, string(ch)
	case ',':
		return comma, string(ch)
	default:
		return corrupted, ""
	}
}

func (s scanner) readIdent() (token, string) {
	defer s.buf.Reset()
	for {
		ch := s.read()
		if ch == 0 {
			break
		}
		if !isLetter(ch) {
			s.unread()
			break
		}
		s.buf.WriteRune(ch)
	}
	literal := s.buf.String()
	switch {
	case strings.HasSuffix(literal, "mul"):
		return mul, literal
	default:
		return corrupted, literal
	}
}

func (s scanner) readNumber() (token, string) {
	defer s.buf.Reset()
	for {
		ch := s.read()
		if ch == 0 {
			break
		}
		if !isNumber(ch) {
			s.unread()
			break
		}
		s.buf.WriteRune(ch)
	}
	return number, s.buf.String()
}

func (s scanner) read() rune {
	r, _, err := s.reader.ReadRune()
	if err != nil {
		return 0
	}
	return r
}

func (s scanner) unread() {
	if err := s.reader.UnreadRune(); err != nil {
		panic(err)
	}
}

func isLetter(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}
