package day16

import (
	"bufio"
	"io"
	"slices"
)

func LowestScore(input io.Reader) int {
	walls := make(map[point]bool)
	var start, end point
	scanner := bufio.NewScanner(input)
	for y := 0; scanner.Scan(); y++ {
		for x, char := range scanner.Text() {
			switch char {
			case '#':
				walls[point{x: x, y: y}] = true
			case 'S':
				start = point{x: x, y: y}
			case 'E':
				end = point{x: x, y: y}
			}
		}
	}

	directions := []direction{directionUp, directionRight, directionDown, directionLeft}
	turnRight := func(dir direction) direction {
		return directions[(slices.Index(directions, dir)+1)%len(directions)]
	}
	turnLeft := func(dir direction) direction {
		return directions[(slices.Index(directions, dir)-1+len(directions))%len(directions)]
	}

	queue := []struct {
		point point
		dir   direction
	}{
		{point: start, dir: directionRight},
	}
	lowestScores := make(map[point]int)
	lowestScores[start] = 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.point == end {
			continue
		}

		for _, step := range []struct {
			dir           direction
			scoreIncrease int
		}{
			{dir: curr.dir, scoreIncrease: 1},
			{dir: turnRight(curr.dir), scoreIncrease: 1001},
			{dir: turnLeft(curr.dir), scoreIncrease: 1001},
		} {
			nextPoint := translate(curr.point, step.dir)
			if walls[nextPoint] {
				continue
			}
			score := lowestScores[curr.point] + step.scoreIncrease
			if lowest, exists := lowestScores[nextPoint]; !exists || score < lowest {
				lowestScores[nextPoint] = score
				queue = append(queue, struct {
					point point
					dir   direction
				}{point: nextPoint, dir: step.dir})
			}
		}
	}

	return lowestScores[end]
}

type point struct{ x, y int }

func translate(p point, dir direction) point {
	switch dir {
	case directionUp:
		return point{x: p.x, y: p.y - 1}
	case directionDown:
		return point{x: p.x, y: p.y + 1}
	case directionLeft:
		return point{x: p.x - 1, y: p.y}
	case directionRight:
		return point{x: p.x + 1, y: p.y}
	}
	panic("invalid direction")
}

type direction int

const (
	directionUp direction = iota
	directionDown
	directionLeft
	directionRight
)
