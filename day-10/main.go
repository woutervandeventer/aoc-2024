package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	switch os.Args[1] {
	case "1":
		fmt.Println(scores(os.Stdin))
	case "2":
		fmt.Println(ratings(os.Stdin))
	}
}

func scores(input io.Reader) (total int) {
	m := readTopographicMap(input)
	heads := getTrailheads(m)

	for _, h := range heads {
		peaks := make(map[position]struct{})
		m.findPeaks(h, func(peak position) { peaks[peak] = struct{}{} })
		total += len(peaks)
	}

	return total
}

func ratings(input io.Reader) (sum int) {
	m := readTopographicMap(input)
	heads := getTrailheads(m)

	for _, h := range heads {
		m.findPeaks(h, func(position) { sum++ })
	}

	return sum
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

func (m topographicMap) findPeaks(start position, peakAction func(peak position)) {
	var climbr func(position)
	climbr = func(p position) {
		height := m.heightAtPos(p)
		if height == 9 {
			peakAction(p)
			return
		}
		for _, next := range []position{
			{x: p.x - 1, y: p.y},
			{x: p.x + 1, y: p.y},
			{x: p.x, y: p.y - 1},
			{x: p.x, y: p.y + 1},
		} {
			if m.heightAtPos(next) == height+1 {
				climbr(next)
			}
		}
	}
	climbr(start)
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
