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

type guard struct {
	x, y int
	dir  direction
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
	return fmt.Sprintf("guard at x: %d, y: %d, facing %s", g.x, g.y, g.dir)
}

type direction byte

func (d direction) String() string {
	return string(d)
}

func main() {
	fmt.Println(distinctPositions(os.Stdin))
}

func distinctPositions(input io.Reader) (count int) {
	var m positionMap
	var g guard
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		// scanner.Bytes() does not allocate, which lead to a nasty bug!
		m = append(m, append([]byte{}, scanner.Bytes()...))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// find the guard
Rows:
	for y, row := range m {
		for x, pos := range row {
			if dir, ok := isGuard(pos); ok {
				g.x = x
				g.y = y
				g.dir = dir
				break Rows
			}
		}
	}

	// Walk the guard off the map and keep track of visited positions
	visited := make(map[int]map[int]bool)
	for guardOnTheMap(g, m) {
		v, ok := visited[g.x]
		if !ok {
			v = make(map[int]bool)
		}
		v[g.y] = true
		visited[g.x] = v
		g.walk(m)
	}
	for _, v := range visited {
		count += len(v)
	}

	return count
}

func isGuard(b byte) (direction, bool) {
	if d := direction(b); d == directionUp || d == directionRight || d == directionDown || d == directionLeft {
		return d, true
	}
	return 0, false
}

func guardOnTheMap(g guard, m positionMap) bool {
	return g.x >= 0 && g.x < len(m[g.x]) && g.y >= 0 && g.y < len(m)
}
