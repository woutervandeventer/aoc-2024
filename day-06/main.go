package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

const (
	obstacle = '#'

	directionUp    direction = '^'
	directionRight direction = '>'
	directionDown  direction = 'v'
	directionLeft  direction = '<'
)

type direction byte

func main() {
	switch os.Args[1] {
	case "1":
		fmt.Println(distinctPositions(os.Stdin))
	case "2":
		fmt.Println(obstructionPositions(os.Stdin))
	}
}

func distinctPositions(input io.Reader) int {
	m := newPositionMap(input)
	g := findGuard(m)

	// Walk the guard off the map and keep track of visited positions
	visited := make(map[position]bool)
	for m.withinBounds(g.position) {
		visited[g.position] = true
		g.walk(m)
	}

	return len(visited)
}

func obstructionPositions(input io.Reader) (count int) {
	m := newPositionMap(input)
	g := findGuard(m)

	var temp positionMap
	for _, row := range m {
		temp = append(temp, append([]byte{}, row...))
	}

	visited := make(map[position]bool)
	for m.withinBounds(g.position) {
		m.lookAhead(g.position, turnRight(g.dir), func(curr, next position) {
			if m.charAt(next) == obstacle && visited[curr] {
				if infront := nextPosition(g.position, g.dir); m.withinBounds(infront) && m.charAt(infront) != obstacle {
					count++
					temp[infront.y][infront.x] = '0'
				}
			}
		})
		visited[g.position] = true
		g.walk(m)
	}

	fmt.Println(temp)

	return count
}

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

func (m positionMap) String() string {
	return string(bytes.Join(m, []byte("\n")))
}

func (m positionMap) charAt(p position) byte {
	return m[p.y][p.x]
}

func (m positionMap) lookAhead(curr position, dir direction, fn func(curr, next position)) {
	for next := nextPosition(curr, dir); m.withinBounds(next); curr, next = next, nextPosition(next, dir) {
		fn(curr, next)
	}
}

func (m positionMap) withinBounds(p position) bool {
	return p.x >= 0 && p.x < len(m[0]) && p.y >= 0 && p.y < len(m)
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
		g.dir = turnRight(g.dir)
		g.walk(m)
		return
	}
	g.position = nextPosition(g.position, g.dir)
}

func nextPosition(p position, dir direction) position {
	switch dir {
	case directionUp:
		p.y--
	case directionRight:
		p.x++
	case directionDown:
		p.y++
	case directionLeft:
		p.x--
	}
	return p
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

func (g *guard) String() string {
	return fmt.Sprintf("guard at %s, facing %s", g.position, g.dir)
}

func turnRight(dir direction) direction {
	switch dir {
	case directionUp:
		return directionRight
	case directionRight:
		return directionDown
	case directionDown:
		return directionLeft
	case directionLeft:
		return directionUp
	default:
		panic("invalid dir")
	}
}

func (d direction) String() string {
	return string(d)
}
