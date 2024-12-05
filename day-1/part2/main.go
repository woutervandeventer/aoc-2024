package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	f, err := os.Open("./day-1/input.txt")
	if err != nil {
		return err
	}
	defer func() { err = errors.Join(err, f.Close()) }()

	var left []int
	right := make(map[int]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l, r, ok := strings.Cut(scanner.Text(), "   ")
		if !ok {
			return errors.New("line does not contain \"   \"")
		}
		lint, err := strconv.Atoi(l)
		if err != nil {
			return err
		}
		left = append(left, lint)
		rint, err := strconv.Atoi(r)
		if err != nil {
			return err
		}
		right[rint] += 1
	}

	var sim int
	for _, lval := range left {
		sim += lval * right[lval]
	}
	fmt.Printf("similarity score: %d\n", sim)

	return nil
}
