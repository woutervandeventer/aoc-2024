package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(fewestTokens(os.Stdin))
}

func fewestTokens(input io.Reader) (tokens int) {
	mm := readMachines(input)
	for _, m := range mm {
		tokens += m.fewestTokens()
	}
	return tokens
}

type machine struct {
	buttonA, buttonB movement
	prize            position
}

func (m machine) String() string {
	format := `Button A: X+%d, Y+%d
Button B: X+%d, Y+%d
Prize: X=%d, Y=%d`
	return fmt.Sprintf(format, m.buttonA.x, m.buttonA.y, m.buttonB.x, m.buttonB.y, m.prize.x, m.prize.y)
}

func (m machine) fewestTokens() int {
	for a := 0; true; a++ {
		remainingx := m.prize.x - m.buttonA.x*a
		remainingy := m.prize.y - m.buttonA.y*a
		if remainingx < 0 || remainingy < 0 {
			return 0
		}
		fmt.Println(remainingx, remainingy)
		divideByWholeNumbers := remainingx%m.buttonB.x == 0 && remainingy%m.buttonB.y == 0
		if b := remainingx / m.buttonB.x; b == remainingy/m.buttonB.y && divideByWholeNumbers {
			return a*3 + b*1
		}
	}
	return 0
}

type movement struct{ x, y int }

type position struct{ x, y int }

func readMachines(r io.Reader) (mm []machine) {
	scanner := bufio.NewScanner(r)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
			return i + 2, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	})
	for scanner.Scan() {
		a, rest, _ := strings.Cut(strings.TrimSpace(scanner.Text()), "\n")
		b, prize, _ := strings.Cut(rest, "\n")
		var m machine
		m.buttonA.x, _ = strconv.Atoi(a[strings.Index(a, "X+")+1 : strings.Index(a, ",")])
		m.buttonA.y, _ = strconv.Atoi(a[strings.Index(a, "Y+")+1:])
		m.buttonB.x, _ = strconv.Atoi(b[strings.Index(b, "X+")+1 : strings.Index(b, ",")])
		m.buttonB.y, _ = strconv.Atoi(b[strings.Index(b, "Y+")+1:])
		m.prize.x, _ = strconv.Atoi(prize[strings.Index(prize, "X=")+2 : strings.Index(prize, ",")])
		m.prize.y, _ = strconv.Atoi(prize[strings.Index(prize, "Y=")+2:])
		const correction = 10000000000000
		m.prize.x += correction
		m.prize.y += correction
		mm = append(mm, m)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return mm
}
