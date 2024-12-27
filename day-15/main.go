package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

func main() {
	fmt.Println(solution(os.Stdin))
}

func solution(input io.Reader) (result int) {
	wh, moves := readInput(input)

	for _, dir := range moves {
		switch wh.objectAt(wh.robot.translate(dir)) {
		case objectWall:
			// Do nothing.
		case objectBox:
			wh.pushBox(dir)
		default:
			wh.moveRobot(dir)
		}
		// fmt.Println(wh)
		// time.Sleep(time.Millisecond * 500)
	}

	fmt.Println(wh)

	for p, o := range wh.objects {
		if o != objectBox {
			continue
		}
		result += p.y*100 + p.x
	}

	return result
}

const (
	objectWall object = '#'
	objectBox  object = 'O'
)

const (
	up direction = iota
	right
	down
	left
)

type warehouse struct {
	objects map[point]object
	robot   point

	width, height int
	canvasOnce    sync.Once
	canvas        [][]rune
}

func (w *warehouse) pushBox(dir direction) {
	var canMove bool
	var boxes []point

	var fn func(point)
	fn = func(p point) {
		obj := w.objectAt(p)
		if obj == objectWall {
			return
		}
		if obj == 0 {
			canMove = true
			return
		}
		if obj == objectBox {
			boxes = append(boxes, p)
			fn(p.translate(dir))
		}
	}
	fn(w.robot.translate(dir))

	if !canMove {
		return
	}
	for i := len(boxes) - 1; i >= 0; i-- {
		box := boxes[i]
		delete(w.objects, box)
		w.objects[box.translate(dir)] = objectBox
	}
	w.moveRobot(dir)
}

func (w *warehouse) objectAt(p point) object {
	return w.objects[p]
}

func (w *warehouse) moveRobot(d direction) {
	switch d {
	case up:
		w.robot.y--
	case down:
		w.robot.y++
	case left:
		w.robot.x--
	case right:
		w.robot.x++
	}
}

func (w *warehouse) String() string {
	w.canvasOnce.Do(func() {
		for p := range w.objects {
			w.width = max(w.width, p.x+1)
			w.height = max(w.height, p.y+1)
		}
		w.canvas = make([][]rune, w.height)
		for i := range w.canvas {
			row := make([]rune, w.width)
			for i := range row {
				row[i] = '.'
			}
			w.canvas[i] = row
		}
	})
	whmap := make([][]rune, len(w.canvas))
	for i, row := range w.canvas {
		whmap[i] = make([]rune, len(row))
		copy(whmap[i], row)
	}
	for p, o := range w.objects {
		whmap[p.y][p.x] = rune(o)
	}
	whmap[w.robot.y][w.robot.x] = '@'
	var b strings.Builder
	for _, row := range whmap {
		b.WriteString(string(row) + "\n")
	}
	return b.String()
}

type point struct{ x, y int }

func (p point) translate(d direction) point {
	switch d {
	case up:
		return point{x: p.x, y: p.y - 1}
	case down:
		return point{x: p.x, y: p.y + 1}
	case left:
		return point{x: p.x - 1, y: p.y}
	case right:
		return point{x: p.x + 1, y: p.y}
	default:
		panic("invalid movement")
	}
}

type object rune

type direction int

func readInput(input io.Reader) (*warehouse, []direction) {
	scanner := bufio.NewScanner(input)
	wh := warehouse{
		objects: make(map[point]object),
	}
	for y := 0; scanner.Scan() && scanner.Text() != ""; y++ {
		for x, r := range scanner.Text() {
			p := point{x: x, y: y}
			switch r {
			case '#':
				wh.objects[p] = objectWall
			case 'O':
				wh.objects[p] = objectBox
			case '@':
				wh.robot = p
			}
		}
	}
	var moves []direction
	for scanner.Scan() {
		for _, r := range scanner.Text() {
			switch r {
			case '^':
				moves = append(moves, up)
			case '>':
				moves = append(moves, right)
			case '<':
				moves = append(moves, left)
			case 'v':
				moves = append(moves, down)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return &wh, moves
}
