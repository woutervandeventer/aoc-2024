package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	day1 "github.com/woutervandeventer/aoc-2024/day-01"
	day2 "github.com/woutervandeventer/aoc-2024/day-02"
	day3 "github.com/woutervandeventer/aoc-2024/day-03"
	day4 "github.com/woutervandeventer/aoc-2024/day-04"
)

func main() {
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var part int
	if len(os.Args) >= 3 {
		part, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
	}
	input := os.Stdin
	switch day {
	case 1:
		switch part {
		case 1:
			fmt.Println(day1.TotalDistance(day1.ReadLists(os.Stdin)))
		case 2:
			fmt.Println(day1.SimilarityScore(day1.ReadLists(os.Stdin)))
		}
	case 2:
		fmt.Println(day2.SafeReportsWithDampener(input))
	case 3:
		fmt.Println(day3.AddMuls(os.Stdin))
	case 4:
		switch part {
		case 1:
			fmt.Println(day4.CountXmas(os.Stdin))
		case 2:
			fmt.Println(day4.CountXMas(os.Stdin))
		}
	case 5:
	case 6:
	case 7:
	case 8:
	case 9:
	case 10:
	case 11:
	case 12:
	case 13:
	case 14:
	case 15:
	}
}
