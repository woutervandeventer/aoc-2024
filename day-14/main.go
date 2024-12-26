package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(safetyFactor(os.Stdin, 101, 103))
}

func safetyFactor(input io.Reader, width, height int) int {
	rr := readRobots(input)
	finalPositions := make(map[point]int)
	for _, r := range rr {
		p := moveRobot(r, width, height, 100)
		finalPositions[p]++
	}
	var leftUpper, rightUpper, leftLower, rightLower int
	for p, n := range finalPositions {
		switch {
		case p.x < width/2 && p.y < height/2:
			leftUpper += n
		case p.x > width/2 && p.y < height/2:
			rightUpper += n
		case p.x < width/2 && p.y > height/2:
			leftLower += n
		case p.x > width/2 && p.y > height/2:
			rightLower += n
		}
	}
	return leftUpper * rightUpper * leftLower * rightLower
}

func moveRobot(r robot, width, height, seconds int) point {
	for range seconds {
		x := r.p.x + r.v.x
		if x < 0 {
			x += width
		}
		if x >= width {
			x -= width
		}
		y := r.p.y + r.v.y
		if y < 0 {
			y += height
		}
		if y >= height {
			y -= height
		}
		r.p.x, r.p.y = x, y
	}
	return r.p
}

func readRobots(r io.Reader) (rr []robot) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		rr = append(rr, parseRobot(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return rr
}

type point struct{ x, y int }

type vector2 struct{ x, y int }

type robot struct {
	p point
	v vector2
}

func (r robot) String() string {
	return fmt.Sprintf("p=%d,%d v=%d,%d", r.p.x, r.p.y, r.v.x, r.v.y)
}

// p=0,4 v=3,-3
func parseRobot(s string) robot {
	var r robot
	pstr, vstr, _ := strings.Cut(s, " ")
	r.p.x, _ = strconv.Atoi(pstr[2:strings.Index(pstr, ",")])
	r.p.y, _ = strconv.Atoi(pstr[strings.Index(pstr, ",")+1:])
	r.v.x, _ = strconv.Atoi(vstr[2:strings.Index(vstr, ",")])
	r.v.y, _ = strconv.Atoi(vstr[strings.Index(vstr, ",")+1:])
	return r
}
