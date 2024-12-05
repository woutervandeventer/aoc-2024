package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
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

	var left, right []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		a, b, ok := strings.Cut(scanner.Text(), "   ")
		if !ok {
			return errors.New("line does not contain \"   \"")
		}
		aint, err := strconv.Atoi(a)
		if err != nil {
			return fmt.Errorf("convert a to int: %w", err)
		}
		left = append(left, aint)
		bint, err := strconv.Atoi(b)
		if err != nil {
			return fmt.Errorf("convert b to int: %w", err)
		}
		right = append(right, bint)
	}

	slices.Sort(left)
	slices.Sort(right)

	var diff int
	for i, leftv := range left {
		rightv := right[i]
		if leftv > rightv {
			diff += leftv - rightv
		} else {
			diff += rightv - leftv
		}
	}
	fmt.Printf("diff: %d\n", diff)

	return nil
}
