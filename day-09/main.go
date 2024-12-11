package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
)

const freeSpace = -1

func main() {
	fmt.Printf("%d\n", checksum(os.Stdin))
}

func checksum(input io.Reader) int {
	m := newDiskmap(input)
	blocks := blocks(m)
	compacted := moveBlocks(blocks)
	return calculateChecksum(compacted)
}

func newDiskmap(input io.Reader) []int {
	b, err := io.ReadAll(input)
	if err != nil {
		panic(err)
	}
	var m []int
	for _, char := range bytes.TrimSpace(b) {
		n, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		m = append(m, n)
	}
	return m
}

func blocks(m []int) []int {
	result := make([]int, 0, len(m))
	var blockID int
	for i, n := range m {
		isBlock := i%2 == 0
		for range n {
			if isBlock {
				result = append(result, blockID)
			} else {
				result = append(result, freeSpace)
			}
		}
		if isBlock {
			blockID++
		}
	}
	return result
}

func moveBlocks(blocks []int) []int {
	for gapsRemaining(blocks) {
		last := lastIndexFunc(blocks, func(n int) bool { return n != freeSpace })
		free := slices.Index(blocks, freeSpace)
		blocks[last], blocks[free] = blocks[free], blocks[last]
	}
	return blocks
}

func gapsRemaining(blocks []int) bool {
	if i := slices.Index(blocks, freeSpace); i >= 0 {
		for _, r := range blocks[i+1:] {
			if r != freeSpace {
				return true
			}
		}
	}
	return false
}

func lastIndexFunc(s []int, fn func(n int) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if fn(s[i]) {
			return i
		}
	}
	return -1
}

func calculateChecksum(compacted []int) (sum int) {
	for i, n := range compacted[:slices.Index(compacted, freeSpace)] {
		sum += i * n
	}
	return sum
}
