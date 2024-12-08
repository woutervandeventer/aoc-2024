package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	withConcat := flag.Bool("concat", false, "With concat operator")
	flag.Parse()
	ops := []func(int, int) int{add, mul}
	if *withConcat {
		ops = append(ops, concat)
	}
	fmt.Println(totalCalibrationResult(os.Stdin, ops...))
}

type equation struct {
	result   int
	operands []int
}

func (e equation) String() string {
	return fmt.Sprintf("%d: %v", e.result, e.operands)
}

func totalCalibrationResult(input io.Reader, ops ...func(int, int) int) (count int) {
	for _, eq := range readEquations(input) {
		if findTarget(eq.result, eq.operands, ops...) {
			count += eq.result
		}
	}
	return count
}

func findTarget(target int, operands []int, ops ...func(int, int) int) bool {
	if len(operands) == 1 {
		return operands[0] == target
	}
	for _, op := range ops {
		if findTarget(target, append([]int{op(operands[0], operands[1])}, operands[2:]...), ops...) {
			return true
		}
	}
	return false
}

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func concat(a, b int) int {
	result, err := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	if err != nil {
		panic(err)
	}
	return result
}

func readEquations(input io.Reader) (eqs []equation) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var eq equation
		line := scanner.Text()
		bef, aft, ok := strings.Cut(line, ": ")
		if !ok {
			panic("separator \": \" not found in line " + line)
		}
		result, err := strconv.Atoi(bef)
		if err != nil {
			panic(err)
		}
		eq.result = result
		for _, s := range strings.Split(aft, " ") {
			op, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			eq.operands = append(eq.operands, op)
		}
		eqs = append(eqs, eq)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return eqs
}
