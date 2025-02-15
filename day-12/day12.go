package day12

import (
	"bufio"
	"io"
)

func TotalPrice(input io.Reader) (price int) {
	var regionMap [][]rune
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		regionMap = append(regionMap, []rune(scanner.Text()))
	}

	type point struct{ x, y int }
	visited := make(map[point]bool)

	for y, row := range regionMap {
		for x, plant := range row {
			if visited[point{x: x, y: y}] {
				continue
			}
			var area, perimeter int

			for queue := []point{{x: x, y: y}}; len(queue) > 0; {
				curr := queue[0]
				queue = queue[1:]
				if visited[curr] {
					continue
				}
				visited[curr] = true
				area++

				for _, neighbour := range []point{
					{x: curr.x + 1, y: curr.y},
					{x: curr.x - 1, y: curr.y},
					{x: curr.x, y: curr.y + 1},
					{x: curr.x, y: curr.y - 1},
				} {
					if neighbour.x < 0 || neighbour.x >= len(row) || neighbour.y < 0 || neighbour.y >= len(regionMap) {
						perimeter++
						continue
					}
					if regionMap[neighbour.y][neighbour.x] == plant {
						queue = append(queue, neighbour)
					} else {
						perimeter++
					}
				}
			}

			price += area * perimeter
		}
	}

	return price
}
