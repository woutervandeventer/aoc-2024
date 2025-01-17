package day6

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

func DistinctGuardPositions(input io.Reader) int {
	lab := readLab(input)
	visited := make(map[point]struct{})
	for lab.withinBounds(lab.guard.position) {
		visited[lab.guard.position] = struct{}{}
		lab.moveGuard()
	}
	return len(visited)
}

func ObstructionPositions(input io.Reader) (count int) {
	lab := readLab(input)
	visitedDirections := make(map[point]direction)
	lab.walkGuard(func(g guard) {
		visitedDirections[g.position] = visitedDirections[g.position] | g.direction
		peekdir := turnRight(g.direction)
		var peek func(curr, next point)
		peek = func(curr, next point) {
			if !lab.withinBounds(curr) {
				return
			}
			if visitedDirections[curr]&turnRight(peekdir) != 0 && lab.isObstruction(next) {
				count++
				return
			}
			peek(next, translate(next, peekdir))
		}
		peek(g.position, translate(g.position, peekdir))
	})
	return count
}

type lab struct {
	obstructions  map[point]struct{}
	guard         guard
	width, height int
}

func readLab(r io.Reader) *lab {
	lab := lab{
		obstructions: make(map[point]struct{}),
	}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		if lab.width == 0 {
			lab.width = utf8.RuneCount(line)
		}
		lab.height++
		for x, ch := range string(line) {
			p := point{x: x, y: lab.height - 1}
			switch ch {
			case '#':
				lab.obstructions[p] = struct{}{}
			case '^':
				lab.guard = guard{
					position:  p,
					direction: directionUp,
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return &lab
}

// walkGuard calls fn until the guard has walked off the map.
func (l *lab) walkGuard(fn func(guard)) {
	for l.withinBounds(l.guard.position) {
		fn(l.guard)
		l.moveGuard()
	}
}

func (l *lab) moveGuard() {
	peek := translate(l.guard.position, l.guard.direction)
	switch l.isObstruction(peek) {
	case true:
		l.guard.direction = turnRight(l.guard.direction)
	case false:
		l.guard.position = peek
	}
}

func (l *lab) isObstruction(pt point) bool {
	_, ok := l.obstructions[pt]
	return ok
}

func (l *lab) withinBounds(pt point) bool {
	return 0 <= pt.x && pt.x < l.width && 0 <= pt.y && pt.y < l.height
}

func (l *lab) draw() {
	canvas := make([][]rune, l.height)
	for i := range canvas {
		background := make([]rune, l.width)
		for j := range background {
			background[j] = '.'
		}
		canvas[i] = background
	}
	for pt := range l.obstructions {
		canvas[pt.y][pt.x] = '#'
	}
	canvas[l.guard.position.y][l.guard.position.x] = l.guard.direction.rune()
	var b strings.Builder
	for _, row := range canvas {
		b.WriteString(string(row) + "\n")
	}
	fmt.Print(b.String())
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

type direction int

const (
	directionUp direction = 1 << iota
	directionRight
	directionDown
	directionLeft
)

func (d direction) rune() rune {
	switch d {
	case directionDown:
		return 'v'
	case directionLeft:
		return '<'
	case directionRight:
		return '>'
	case directionUp:
		return '^'
	default:
		return 'x'
	}
}

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
