package day18

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func MinimumStepsToExit(gridSize, linesToRead int, bytePositions io.Reader) (steps int) {
	corrupted := make(map[point]bool, linesToRead)
	scanner := bufio.NewScanner(bytePositions)
	for range linesToRead {
		scanner.Scan()
		xstr, ystr, _ := strings.Cut(string(scanner.Bytes()), ",")
		x, _ := strconv.Atoi(xstr)
		y, _ := strconv.Atoi(ystr)
		corrupted[point{x: x, y: y}] = true
	}

	// Recursively solve the maze, keeping track of the current minimal amount of steps
	visited := make(map[point]bool)

	var walk func(p point)
	walk = func(p point) {
		if 0 > p.x || p.x >= gridSize || 0 > p.y || p.y >= gridSize {
			return
		}
		if corrupted[p] {
			return
		}
		if visited[p] {
			return
		}
		if steps != 0 && len(visited) >= steps {
			return
		}
		if p.x == gridSize-1 && p.y == gridSize-1 {
			steps = len(visited)
			return
		}
		visited[p] = true
		defer func() { delete(visited, p) }()
		walk(point{x: p.x, y: p.y + 1})
		walk(point{x: p.x + 1, y: p.y})
		walk(point{x: p.x, y: p.y - 1})
		walk(point{x: p.x - 1, y: p.y})
	}
	walk(point{x: 0, y: 0})

	return steps
}

type point struct{ x, y int }

func draw(gridSize int, corrupted, visited map[point]bool) {
	canvas := make([][]rune, gridSize)
	for i := range gridSize {
		canvas[i] = make([]rune, gridSize)
		for j := range gridSize {
			canvas[i][j] = '.'
		}
	}
	for p := range corrupted {
		canvas[p.y][p.x] = '#'
	}
	for p := range visited {
		canvas[p.y][p.x] = '0'
	}
	for _, line := range canvas {
		fmt.Println(string(line))
	}
}
