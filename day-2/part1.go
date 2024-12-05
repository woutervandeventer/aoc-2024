package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err.Error())
	}
}

func run() error {
	f, err := os.Open("./day-2/reports.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	var safe int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ss := strings.Split(scanner.Text(), " ")
		report := make([]int, len(ss))
		for i, s := range ss {
			if report[i], err = strconv.Atoi(s); err != nil {
				return err
			}
		}
		b := isSafe(report)
		fmt.Printf("%v: %t\n", report, b)
		if b {
			safe++
		}
	}
	fmt.Println(safe)
	return nil
}

func isSafe(report []int) bool {
	return isSafeFn(0, report)
}

func isSafeFn(retries int, report []int) bool {
	if retries > 1 {
		return false
	}
	var asc bool
	for i, j := 0, 1; i < len(report)-1; i, j = i+1, j+1 {
		left, right := report[i], report[j]
		currentlyAsc := right > left
		if i == 0 {
			asc = currentlyAsc
		}
		woleft := append(append([]int{}, report[:i]...), report[i+1:]...)
		woright := append(append([]int{}, report[:j]...), report[j+1:]...)
		if left == right || diff(left, right) > 3 {
			return isSafeFn(retries+1, woleft) || isSafeFn(retries+1, woright)
		}
		if asc != currentlyAsc {
			woprev := append(append([]int{}, report[:i-1]...), report[i:]...)
			return isSafeFn(retries+1, woprev) || isSafeFn(retries+1, woleft) || isSafeFn(retries+1, woright)
		}
	}
	return true
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
