package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(scores(os.Stdin))
}

func scores(input io.Reader) (total int) {
	// m := newTopographicMap(input)
	// trailheads := m.trailheads()
	// for _, th := range trailheads {
	// 	total += m.scores(th)
	// }
	_ = input
	return total
}

// type topographicMap [][]int

// func newTopographicMap(r io.Reader) (m topographicMap) {
// 	scanner := bufio.NewScanner(r)
// 	for scanner.Scan() {
// 		var row []int
// 		for _, b := range scanner.Bytes() {
// 			n, err := strconv.Atoi(string(b))
// 			if err != nil {
// 				panic(err)
// 			}
// 			row = append(row, n)
// 		}
// 		m = append(m, row)
// 	}
// 	return m
// }

// func (m topographicMap) String() string {
// 	var b strings.Builder
// 	for _, row := range m {
// 		b.WriteString(fmt.Sprintln(row))
// 	}
// 	return b.String()
// }

// func (m topographicMap) trailheads() []position {
// 	var result []position
// 	for y, row := range m {
// 		for x, n := range row {
// 			if n == 0 {
// 				result = append(result, position{x: x, y: y})
// 			}
// 		}
// 	}
// 	return result
// }

// func (m topographicMap) height(p position) int {
// 	// Check bounds
// 	if p.x < 0 || p.y < 0 || p.x >= len(m[p.y]) || p.y >= len(m) {
// 		return -1
// 	}
// 	return m[p.y][p.x]
// }

// func (m topographicMap) scores(p position) int {
// 	height := m.height(p)
// 	if height == 9 {
// 		return 1
// 	}
// 	up := position{x: p.x, y: p.y - 1}
// 	right := position{x: p.x + 1, y: p.y}
// 	down := position{x: p.x, y: p.y + 1}
// 	left := position{x: p.x - 1, y: p.y}
// 	return scores(up)
// }

// func (m topographicMap) tryClimb(curr, next position) []position {
// 	currHeight, nextHeight := m.height(curr), m.height(next)
// 	if nextHeight != currHeight+1 {
// 		return nil
// 	}
// 	if nextHeight == 9 {
// 		return []position{next}
// 	}
// 	up := position{x: next.x, y: next.y - 1}
// 	right := position{x: next.x + 1, y: next.y}
// 	down := position{x: next.x, y: next.y + 1}
// 	left := position{x: next.x - 1, y: next.y}
// 	return append(
// 		append(
// 			append(
// 				append([]position{},
// 					m.tryClimb(next, up)...),
// 				m.tryClimb(next, right)...),
// 			m.tryClimb(next, down)...),
// 		m.tryClimb(next, left)...)
// }

// type position struct{ x, y int }

// func (p position) String() string {
// 	return fmt.Sprintf("x: %d, y: %d", p.x, p.y)
// }
