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
	for _, antennas := range antennasByFreq {
		for i := range antennas {
			for j := i + 1; j < len(antennas); j++ {
				a, b := antennas[i], antennas[j]
				nextRec(a, b, func(p position) bool {
					within := m.withinBounds(p)
					if within {
						antinodes[p] = true
					}
					return within
				})
				nextRec(b, a, func(p position) bool {
					within := m.withinBounds(p)
					if within {
						antinodes[p] = true
					}
					return within
				})
				antinodes[a], antinodes[b] = true, true
			}
		}
	}
	for p := range antinodes {
		if char := m[p.y][p.x]; char == '.' {
			m[p.y][p.x] = '#'
		}
	}
	return len(antinodes)
}

func nextRec(a, b position, cont func(position) bool) {
	if next := nextAntiNode(a, b); cont(next) {
		nextRec(b, next, cont)
	}
}

func nextAntiNode(a, b position) position {
	return position{x: b.x + (b.x - a.x), y: b.y + (b.y - a.y)}
}

type cityMap [][]byte

func (m cityMap) withinBounds(p position) bool {
	return p.x >= 0 && p.x < len(m[0]) && p.y >= 0 && p.y < len(m)
}

type position struct{ x, y int }

func (p position) String() string {
	return fmt.Sprintf("x=%d, y=%d", p.x, p.y)
}

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
