package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(stones(os.Stdin, 25))
}

func stones(input io.Reader, blinks int) int {
	_, _ = input, blinks
	stones := readStones(input)
	for range blinks {
		var newStones []stone
		for _, s := range stones {
			newStones = append(newStones, s.transform()...)
		}
		stones = newStones
	}
	return len(stones)
}

type stone int

func readStones(r io.Reader) []stone {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		panic(err)
	}
	var stones []stone
	for _, s := range strings.Split(strings.TrimSpace(buf.String()), " ") {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		stones = append(stones, stone(n))
	}
	return stones
}

func (s stone) transform() []stone {
	if s == 0 {
		return []stone{1}
	}
	if str := strconv.Itoa(int(s)); len(str)%2 == 0 {
		left, err := strconv.Atoi(str[:len(str)/2])
		if err != nil {
			panic(err)
		}
		right, err := strconv.Atoi(str[len(str)/2:])
		if err != nil {
			panic(err)
		}
		return []stone{stone(left), stone(right)}
	}
	return []stone{s * 2024}
}
