package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(stones(readStones(os.Stdin), 75))
}

func stones(stones []stone, blinks int) (count int) {
	type problem struct {
		stone  stone
		blinks int
	}
	solutions := make(map[problem]int)

	var transform func(stone, int) int
	transform = func(s stone, blinks int) (result int) {
		if blinks <= 0 {
			return 1
		}
		if solution, ok := solutions[problem{stone: s, blinks: blinks}]; ok {
			return solution
		}
		defer func() {
			solutions[problem{s, blinks}] = result
		}()
		if s == "0" {
			return transform("1", blinks-1)
		}
		if len(s)%2 == 0 {
			right := s[len(s)/2:]
			for strings.HasPrefix(right, "0") && len(right) > 1 {
				right = strings.TrimPrefix(right, "0")
			}
			return transform(s[:len(s)/2], blinks-1) + transform(right, blinks-1)
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return transform(strconv.Itoa(n*2024), blinks-1)
	}

	for _, s := range stones {
		count += transform(s, blinks)
	}

	return count
}

type stone = string

func readStones(r io.Reader) []stone {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(buf.String()), " ")
}
