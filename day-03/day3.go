package day3

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
)

var mulrexp = regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)

func AddMuls(input string) (result int) {
	for _, s := range mulrexp.FindAllString(input, -1) {
		result += execMulStr(s)
	}
	return result
}

func AddEnabledMuls(input []byte) (result int) {
	enabled := true
	for i, b := range input {
		switch b {
		case 'm':
			if !enabled {
				continue
			}
			distanceToClosingParen := bytes.IndexByte(input[i:], ')')
			if distanceToClosingParen < 0 {
				continue
			}
			if mulstr := input[i : i+distanceToClosingParen+1]; mulrexp.Match(mulstr) {
				result += execMulStr(string(mulstr))
			}
		case 'd':
			switch {
			case bytes.HasPrefix(input[i:], []byte("don't()")):
				enabled = false
			case bytes.HasPrefix(input[i:], []byte("do()")):
				enabled = true
			}
		}
	}
	return result
}

func execMulStr(mul string) int {
	mul = mul[4 : len(mul)-1]
	a, b, _ := strings.Cut(mul, ",")
	aint, _ := strconv.Atoi(a)
	bint, _ := strconv.Atoi(b)
	return aint * bint
}
