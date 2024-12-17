package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(totalPrice(os.Stdin))
}

func totalPrice(input io.Reader) (total int) {
	m := readRegionMap(input)

	visited := make(map[plot]struct{})
	for y, row := range m {
		for x, plant := range row {
			p := plot{x: x, y: y}
			if _, ok := visited[p]; ok {
				continue
			}
			var area, perimeter int
			var visit func(plot)
			visit = func(p plot) {
				if _, ok := visited[p]; ok {
					return
				}
				visited[p] = struct{}{}
				area++
				for _, direction := range []plot{
					{x: p.x + 1, y: p.y},
					{x: p.x - 1, y: p.y},
					{x: p.x, y: p.y + 1},
					{x: p.x, y: p.y - 1},
				} {
					if m.plantAt(direction) == plant {
						visit(direction)
					} else {
						perimeter++
					}
				}
			}
			visit(p)
			total += area * perimeter
		}
	}

	return total
}

type regionMap [][]rune

func readRegionMap(r io.Reader) (m regionMap) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		m = append(m, []rune(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return m
}

func (m regionMap) plantAt(p plot) rune {
	if p.x < 0 || p.y < 0 || p.x >= len(m[0]) || p.y >= len(m) {
		return ' '
	}
	return m[p.y][p.x]
}

type plot struct{ x, y int }
