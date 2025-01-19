package day18

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func MinimumStepsToExit(gridSize, linesToRead int, bytePositions io.Reader) int {
	corrupted := make(map[point]bool, linesToRead)
	scanner := bufio.NewScanner(bytePositions)
	for range linesToRead {
		scanner.Scan()
		xstr, ystr, _ := strings.Cut(string(scanner.Bytes()), ",")
		x, _ := strconv.Atoi(xstr)
		y, _ := strconv.Atoi(ystr)
		corrupted[point{x: x, y: y}] = true
	}

	type step struct {
		pt    point
		steps int
	}
	queue := []step{
		{}, // Zero value resembles the starting position.
	}
	visited := make(map[point]bool)
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, step := range []step{
			{pt: point{x: curr.pt.x + 1, y: curr.pt.y}, steps: curr.steps + 1},
			{pt: point{x: curr.pt.x, y: curr.pt.y + 1}, steps: curr.steps + 1},
			{pt: point{x: curr.pt.x - 1, y: curr.pt.y}, steps: curr.steps + 1},
			{pt: point{x: curr.pt.x, y: curr.pt.y - 1}, steps: curr.steps + 1},
		} {
			if 0 > step.pt.x || step.pt.x > gridSize || 0 > step.pt.y || step.pt.y > gridSize ||
				corrupted[step.pt] ||
				visited[step.pt] {
				continue
			}
			if step.pt.x == gridSize && step.pt.y == gridSize {
				return step.steps
			}
			visited[step.pt] = true
			queue = append(queue, step)
		}
	}

	return 0
}

type point struct{ x, y int }
