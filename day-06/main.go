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
	for ; lab.withinBounds(lab.guard.position); lab.moveGuard() {
		visited[lab.guard.position] = struct{}{}
	}
	return len(visited)
}

func ObstructionPositions(input io.Reader) int {
	l := readLab(input)
	obstructions := make(map[point]struct{})

	for ; l.withinBounds(l.guard.position); l.moveGuard() {
		next := translate(l.guard.position, l.guard.direction)
		if !l.withinBounds(next) || l.isObstruction(next) {
			continue
		}
		if enterLoopWithObstacle(next, l) {
			obstructions[next] = struct{}{}
		}
	}

	return len(obstructions)
}

func enterLoopWithObstacle(obstacle point, l *lab) bool {
	// Temporarily add the obstacle to the lab, to see where it sends the guard.
	// The position of the guard has to be restored at the end of this test.
	l.addObstruction(obstacle)
	snapshot := l.guard
	defer func() {
		l.removeObstruction(obstacle)
		l.guard = snapshot
	}()

	visitedDirections := make(map[point]direction)
	for ; l.withinBounds(l.guard.position); l.moveGuard() {
		// Check if the bit is set for the current direction of the guard. If so,
		// we've entered a loop. If not, set it and continue
		directions := visitedDirections[l.guard.position]
		if directions&l.guard.direction != 0 {
			fmt.Println("obstacle position found:", obstacle)
			l.drawWithPath(visitedDirections, obstacle)
			return true
		}
		visitedDirections[l.guard.position] = directions | l.guard.direction
	}

	// The guard has walked off the map.
	return false
}

type lab struct {
	obstructions map[point]struct{}
	guard        struct {
		position  point
		direction direction
	}
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
				lab.addObstruction(p)
			case '^':
				lab.guard.position = p
				lab.guard.direction = directionUp
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return &lab
}

func (l *lab) addObstruction(p point) {
	l.obstructions[p] = struct{}{}
}

func (l *lab) removeObstruction(p point) {
	delete(l.obstructions, p)
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

func (l *lab) drawWithPath(visitedDirections map[point]direction, obstacle point) {
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
	for pt, dirs := range visitedDirections {
		var horizontal, vertical bool
		if dirs&directionUp != 0 || dirs&directionDown != 0 {
			vertical = true
		}
		if dirs&directionLeft != 0 || dirs&directionRight != 0 {
			horizontal = true
		}
		var ch rune
		switch {
		case horizontal && vertical:
			ch = '+'
		case horizontal:
			ch = '-'
		case vertical:
			ch = '|'
		}
		canvas[pt.y][pt.x] = ch
	}
	canvas[obstacle.y][obstacle.x] = '0'
	canvas[l.guard.position.y][l.guard.position.x] = l.guard.direction.rune()
	var b strings.Builder
	for _, row := range canvas {
		b.WriteString(string(row) + "\n")
	}
	fmt.Print(b.String())
}

type point struct{ x, y int }

func (p point) String() string {
	return fmt.Sprintf("x=%d y=%d", p.x, p.y)
}

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
		return '?'
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
