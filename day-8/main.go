package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(uniqueAntinodes(os.Stdin))
}

func uniqueAntinodes(input io.Reader) int {
	m := newCityMap(input)
	antennasByFreq := make(map[byte][]position)
	for y, row := range m {
		for x, char := range row {
			if char != '.' {
				antennasByFreq[char] = append(antennasByFreq[char], position{x: x, y: y})
			}
		}
	}
	antinodes := make(map[position]bool)
	for _, positions := range antennasByFreq {
		for i, curr := range positions {
			for j := i + 1; j < len(positions); j++ {
				next := positions[j]
				node1 := position{x: curr.x - (next.x - curr.x), y: curr.y - (next.y - curr.y)}
				if m.withinBounds(node1) {
					antinodes[node1] = true
				}
				node2 := position{x: next.x + (next.x - curr.x), y: next.y + (next.y - curr.y)}
				if m.withinBounds(node2) {
					antinodes[node2] = true
				}
			}
		}
	}
	for p := range antinodes {
		m[p.y][p.x] = '#'
	}
	return len(antinodes)
}

type cityMap [][]byte

func (m cityMap) withinBounds(p position) bool {
	return p.x >= 0 && p.x < len(m[0]) && p.y >= 0 && p.y < len(m)
}

type position struct{ x, y int }

func newCityMap(r io.Reader) (m cityMap) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		m = append(m, append([]byte{}, scanner.Bytes()...))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return m
}

func (m cityMap) String() string {
	return string(bytes.Join(m, []byte("\n")))
}
