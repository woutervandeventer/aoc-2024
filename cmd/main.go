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
	day5 "github.com/woutervandeventer/aoc-2024/day-05"
	day6 "github.com/woutervandeventer/aoc-2024/day-06"
	day9 "github.com/woutervandeventer/aoc-2024/day-09"
	day12 "github.com/woutervandeventer/aoc-2024/day-12"
	day16 "github.com/woutervandeventer/aoc-2024/day-16"
	day18 "github.com/woutervandeventer/aoc-2024/day-18"
	day19 "github.com/woutervandeventer/aoc-2024/day-19"
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
			fmt.Println(day1.TotalDistance(day1.ReadLists(input)))
		case 2:
			fmt.Println(day1.SimilarityScore(day1.ReadLists(input)))
		}
	case 2:
		fmt.Println(day2.SafeReportsWithDampener(input))
	case 3:
		fmt.Println(day3.AddMuls(input))
	case 4:
		switch part {
		case 1:
			fmt.Println(day4.CountXmas(input))
		case 2:
			fmt.Println(day4.CountXMas(input))
		}
	case 5:
		fmt.Println(day5.SumMiddlePageNos(input))
	case 6:
		switch part {
		case 1:
			fmt.Println(day6.DistinctGuardPositions(input))
		case 2:
			fmt.Println(day6.ObstructionPositions(input))
		}
	case 9:
		fmt.Println(day9.Checksum(input))
	case 12:
		fmt.Println(day12.TotalPrice(input))
	case 16:
		fmt.Println(day16.LowestScore(input))
	case 18:
		switch part {
		case 1:
			fmt.Println(day18.MinimumStepsToExit(70, 1024, input))
		case 2:
			fmt.Println(day18.BlockingByteCoordinates(70, 1024, input))
		}
	case 19:
		fmt.Println(day19.PossibleTowelCombinations(input))
	}
}
