package day9

import (
	"io"
	"strconv"
)

func Checksum(input io.Reader) (checksum int) {
	const freeSpace = -1

	type file struct{ index, width int }
	var files []file
	var blocks []int

	bytes, _ := io.ReadAll(input)
	for i, b := range bytes {
		n, _ := strconv.Atoi(string(b))
		var block int
		switch {
		case i%2 == 0: // on a file
			files = append(files, file{index: len(blocks), width: n})
			block = i / 2
		default:
			block = freeSpace
		}
		for range n {
			blocks = append(blocks, block)
		}
	}

Files:
	for fileID := len(files) - 1; fileID > 0; fileID-- {
		file := files[fileID]
		for i := 0; i < file.index-file.width; {
			for blocks[i] != freeSpace {
				i++
			}
			for j := i; j < file.index; j++ {
				if blocks[j] != freeSpace {
					i = j
					break
				}
				if j-i+1 == file.width {
					for n := range file.width {
						blocks[i+n], blocks[file.index+n] = blocks[file.index+n], blocks[i+n]
					}
					continue Files
				}
			}
		}
	}

	// Calculate checksum
	for i, block := range blocks {
		if block == freeSpace {
			continue
		}
		checksum += i * block
	}

	return checksum
}
