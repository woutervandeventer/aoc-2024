package day9

import (
	"bytes"
	"io"
	"strconv"
)

const freeSpace = -1

func Checksum(input io.Reader) int {
	blocks := blocks(newDiskmap(input))

	// compact the blocks

	return calculateChecksum(blocks)
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

func blocks(diskmap []int) []int {
	result := make([]int, 0, len(diskmap))
	var blockID int
	for i, n := range diskmap {
		isBlock := i%2 == 0
		for range n {
			switch {
			case isBlock:
				result = append(result, blockID)
			default:
				result = append(result, freeSpace)
			}
		}
		if isBlock {
			blockID++
		}
	}
	return result
}

func calculateChecksum(compacted []int) (sum int) {
	for i, n := range compacted {
		if n != freeSpace {
			sum += i * n
		}
	}
	return sum
}
