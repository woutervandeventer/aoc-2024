package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	fmt.Println(scores(os.Stdin))
}

func scores(input io.Reader) (total int) {
	m := readTopographicMap(input)
	heads := getTrailheads(m)

	for _, h := range heads {
		peaks := make(map[position]bool)
		var findScores func(position)
		findScores = func(p position) {
			height := m.heightAtPos(p)
			if height == 9 {
				peaks[p] = true
				return
			}
			left := position{x: p.x - 1, y: p.y}
			if m.heightAtPos(left) == height+1 {
				findScores(left)
			}
			right := position{x: p.x + 1, y: p.y}
			if m.heightAtPos(right) == height+1 {
				findScores(right)
			}
			up := position{x: p.x, y: p.y - 1}
			if m.heightAtPos(up) == height+1 {
				findScores(up)
			}
			down := position{x: p.x, y: p.y + 1}
			if m.heightAtPos(down) == height+1 {
				findScores(down)
			}
		}
		findScores(h)
		total += len(peaks)
	}

	return total
}

type topographicMap [][]int

func readTopographicMap(r io.Reader) (m topographicMap) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var row []int
		for _, b := range scanner.Bytes() {
			n, err := strconv.Atoi(string(b))
			if err != nil {
				panic(err)
			}
			row = append(row, n)
		}
		m = append(m, row)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return m
}

func (m topographicMap) heightAtPos(p position) int {
	if p.x < 0 || p.y < 0 || p.x >= len(m[0]) || p.y >= len(m) {
		return -1
	}
	return m[p.y][p.x]
}

type position struct{ x, y int }

func getTrailheads(m topographicMap) (heads []position) {
	for y, row := range m {
		for x, height := range row {
			if height == 0 {
				heads = append(heads, position{x: x, y: y})
			}
		}
	}
	return heads
}
