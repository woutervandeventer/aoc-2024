package day18

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func MinimumStepsToExit(gridSize, linesToRead int, input io.Reader) int {
	corrupted := make(map[point]bool, linesToRead)
	scanner := newPointScanner(input)
	for range linesToRead {
		scanner.scan()
		corrupted[scanner.point] = true
	}

	// Do a BFS over the grid.
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

func BlockingByteCoordinates(gridSize, linesToRead int, input io.Reader) string {
	corrupted := make(map[point]bool, linesToRead)
	scanner := newPointScanner(input)
	for range linesToRead {
		scanner.scan()
		corrupted[scanner.point] = true
	}

ScanLoop:
	for scanner.scan() {
		// Drop a byte
		corrupted[scanner.point] = true

		// Check if the exit can still be reached.
		queue := []point{
			{}, // Zero value resembles the starting position.
		}
		visited := make(map[point]bool)
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			for _, pt := range []point{
				{x: curr.x + 1, y: curr.y},
				{x: curr.x, y: curr.y + 1},
				{x: curr.x - 1, y: curr.y},
				{x: curr.x, y: curr.y - 1},
			} {
				if 0 > pt.x || pt.x > gridSize || 0 > pt.y || pt.y > gridSize ||
					corrupted[pt] ||
					visited[pt] {
					continue
				}
				if pt.x == gridSize && pt.y == gridSize {
					continue ScanLoop
				}
				visited[pt] = true
				queue = append(queue, pt)
			}
		}
		// We weren't able to find the exit, so the last byte blocked the way.
		return fmt.Sprintf("%d,%d", scanner.point.x, scanner.point.y)
	}

	return ""
}

type point struct{ x, y int }

type pointScanner struct {
	scanner *bufio.Scanner
	point   point
}

func newPointScanner(r io.Reader) *pointScanner {
	return &pointScanner{scanner: bufio.NewScanner(r)}
}

func (ps *pointScanner) scan() bool {
	if !ps.scanner.Scan() {
		return false
	}
	xstr, ystr, _ := strings.Cut(string(ps.scanner.Bytes()), ",")
	ps.point.x, _ = strconv.Atoi(xstr)
	ps.point.y, _ = strconv.Atoi(ystr)
	return true
}
