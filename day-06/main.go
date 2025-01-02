package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

func main() {
	switch os.Args[1] {
	case "1":
		fmt.Println(distinctPositions(os.Stdin))
	case "2":
		fmt.Println(obstructionPositions(os.Stdin))
	}
}

func distinctPositions(input io.Reader) int {
	lab := readLab(input)
	visited := make(map[point]struct{})
	lab.walkGuard(func(g guard) {
		visited[g.position] = struct{}{}
	})
	return len(visited)
}

func obstructionPositions(input io.Reader) (count int) {
	lab := readLab(input)
	visitedDirections := make(map[point][]direction)
	lab.walkGuard(func(g guard) {
		peekdir := turnRight(g.direction)
		for next := translate(g.position, peekdir); lab.withinBounds(next); next = translate(next, peekdir) {
			if lab.isObstruction(next) {
				break // Blocked by obstruction.
			}
			if slices.Contains(visitedDirections[next], peekdir) {
				count++
				break
			}
		}
		visitedDirections[g.position] = append(visitedDirections[g.position], g.direction)
	})
	return count
}

type lab struct {
	obstructions  map[point]struct{}
	guard         guard
	width, height int
}

func readLab(r io.Reader) lab {
	lab := lab{
		obstructions: make(map[point]struct{}),
	}
	scanner := bufio.NewScanner(r)
	for y := 0; scanner.Scan(); y++ {
		row := scanner.Text()
		for x, r := range row {
			p := point{x: x, y: y}
			switch r {
			case '#':
				lab.obstructions[p] = struct{}{}
			case '^':
				lab.guard = guard{
					position:  p,
					direction: directionUp,
				}
			}
		}
		if lab.width == 0 {
			lab.width = len(row)
		}
		lab.height++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lab
}

// walkGuard calls fn until the guard has walked off the map.
func (l *lab) walkGuard(fn func(guard)) {
	var walkFn func()
	walkFn = func() {
		if l.guard.position.x >= l.width || l.guard.position.y >= l.height {
			return
		}
		fn(l.guard)
		next := translate(l.guard.position, l.guard.direction)
		if l.isObstruction(next) {
			l.guard.direction = turnRight(l.guard.direction)
		} else {
			l.guard.position = next
		}
		walkFn()
	}
	walkFn()
}

func (l *lab) isObstruction(pt point) bool {
	_, ok := l.obstructions[pt]
	return ok
}

func (l *lab) withinBounds(pt point) bool {
	return pt.x >= 0 && pt.x < l.width && pt.y >= 0 && pt.y < l.height
}

func (l *lab) draw() {
	row := make([]rune, l.width)
	for i := range row {
		row[i] = '.'
	}
	canvas := make([][]rune, l.height)
	for i := range canvas {
		canvas[i] = append(make([]rune, 0, len(row)), row...)
	}
	for pt := range l.obstructions {
		canvas[pt.y][pt.x] = '#'
	}
	switch l.guard.direction {
	case directionUp:
		canvas[l.guard.position.y][l.guard.position.x] = '^'
	case directionRight:
		canvas[l.guard.position.y][l.guard.position.x] = '>'
	case directionDown:
		canvas[l.guard.position.y][l.guard.position.x] = 'v'
	case directionLeft:
		canvas[l.guard.position.y][l.guard.position.x] = '<'
	}
	for _, row := range canvas {
		fmt.Println(string(row))
	}
}

type guard struct {
	position  point
	direction direction
}

type point struct{ x, y int }

func translate(pt point, dir direction) point {
	switch dir {
	case directionDown:
		return point{x: pt.x, y: pt.y + 1}
	case directionLeft:
		return point{x: pt.x - 1, y: pt.y}
	case directionRight:
		return point{x: pt.x + 1, y: pt.y}
	case directionUp:
		return point{x: pt.x, y: pt.y - 1}
	default:
		panic("invalid direction")
	}
}

const (
	directionUp direction = iota
	directionRight
	directionDown
	directionLeft
)

type direction int

func turnRight(dir direction) direction {
	switch dir {
	case directionDown:
		return directionLeft
	case directionLeft:
		return directionUp
	case directionRight:
		return directionDown
	case directionUp:
		return directionRight
	default:
		panic("invalid direction")
	}
}
