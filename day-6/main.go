package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	obstacle       = '#'
	directionUp    = direction('^')
	directionRight = direction('>')
	directionDown  = direction('v')
	directionLeft  = direction('<')
)

type positionMap [][]byte

func newPositionMap(r io.Reader) (m positionMap) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		// scanner.Bytes() does not allocate, which lead to a nasty bug!
		m = append(m, append([]byte{}, scanner.Bytes()...))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return m
}

type position struct{ x, y int }

func (p position) String() string { return fmt.Sprintf("x: %d, y: %d", p.x, p.y) }

type guard struct {
	position
	dir direction
}

func findGuard(m positionMap) guard {
	for y, row := range m {
		for x, pos := range row {
			if dir, ok := isGuard(pos); ok {
				return guard{
					position: position{x: x, y: y},
					dir:      dir,
				}
			}
		}
	}
	panic("guard not found")
}

func isGuard(b byte) (direction, bool) {
	if d := direction(b); d == directionUp || d == directionRight || d == directionDown || d == directionLeft {
		return d, true
	}
	return 0, false
}

func (g *guard) walk(m positionMap) {
	if g.peekObstacle(m) {
		g.turnRight()
		g.walk(m)
		return
	}
	switch g.dir {
	case directionUp:
		g.y--
	case directionRight:
		g.x++
	case directionDown:
		g.y++
	case directionLeft:
		g.x--
	}
}

func (g *guard) peekObstacle(m positionMap) bool {
	switch g.dir {
	case directionUp:
		if g.y == 0 {
			return false
		}
		return m[g.y-1][g.x] == obstacle
	case directionRight:
		if g.x == len(m[g.x])-1 {
			return false
		}
		return m[g.y][g.x+1] == obstacle
	case directionDown:
		if g.y == len(m)-1 {
			return false
		}
		return m[g.y+1][g.x] == obstacle
	case directionLeft:
		if g.x == 0 {
			return false
		}
		return m[g.y][g.x-1] == obstacle
	default:
		panic("invalid guard direction: " + string(g.dir))
	}
}

func (g *guard) onTheMap(m positionMap) bool {
	return g.x >= 0 && g.x < len(m[g.x]) && g.y >= 0 && g.y < len(m)
}

func (g *guard) turnRight() {
	switch g.dir {
	case directionUp:
		g.dir = directionRight
	case directionRight:
		g.dir = directionDown
	case directionDown:
		g.dir = directionLeft
	case directionLeft:
		g.dir = directionUp
	}
}

func (g *guard) String() string {
	return fmt.Sprintf("guard at %s, facing %s", g.position, g.dir)
}

type direction byte

func (d direction) String() string {
	return string(d)
}

func main() {
	fmt.Println(distinctPositions(os.Stdin))
}

func distinctPositions(input io.Reader) int {
	m := newPositionMap(input)
	g := findGuard(m)

	// Walk the guard off the map and keep track of visited positions
	visited := make(map[position]bool)
	for g.onTheMap(m) {
		visited[g.position] = true
		g.walk(m)
	}

	return len(visited)
}

func obstructionPositions(input io.Reader) (count int) {
	m := newPositionMap(input)
	g := findGuard(m)

	visited := make(map[position]direction)
	for g.onTheMap(m) {
		// check if we've been here before, and if the direction was our direction turned to the right
		if dir, seen := visited[g.position]; seen {
			fmt.Println("we have been at", g.position)
			switch g.dir {
			case directionUp:
				if dir == directionRight {
					count++
				}
			case directionRight:
				if dir == directionDown {
					count++
				}
			case directionDown:
				if dir == directionLeft {
					count++
				}
			case directionLeft:
				if dir == directionUp {
					count++
				}
			}
		} else {
			visited[g.position] = g.dir
		}
		g.walk(m)
	}

	return count
}
