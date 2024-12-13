package main

import (
	"flag"
	"fmt"

	"github.com/mbe81/advent-of-code-2024/days/day01"
	"github.com/mbe81/advent-of-code-2024/days/day02"
	"github.com/mbe81/advent-of-code-2024/days/day03"
	"github.com/mbe81/advent-of-code-2024/days/day04"
	"github.com/mbe81/advent-of-code-2024/days/day05"
	"github.com/mbe81/advent-of-code-2024/days/day06"
	"github.com/mbe81/advent-of-code-2024/days/day07"
	"github.com/mbe81/advent-of-code-2024/days/day08"
	"github.com/mbe81/advent-of-code-2024/days/day09"
	"github.com/mbe81/advent-of-code-2024/days/day10"
	"github.com/mbe81/advent-of-code-2024/days/day11"
	"github.com/mbe81/advent-of-code-2024/days/day12"
	"github.com/mbe81/advent-of-code-2024/days/day13"
)

func main() {
	day := flag.Int("day", 0, "Select the day to run")
	part := flag.Int("part", 0, "Select the part to run (optional)")
	filename := flag.String("input", "input.txt", "Provide the input file to use (optional)")
	flag.Parse()

	if *day == 0 || *filename == "" {
		fmt.Println("Invalid flags. Please use the following flags:\n")
		flag.PrintDefaults()
		fmt.Println("\nExample: go run ./cmd/aoc -day 1 -part 1 -input input.txt")
		return
	}

	switch *day {
	case 1:
		day01.Run(*part, *filename)
	case 2:
		day02.Run(*part, *filename)
	case 3:
		day03.Run(*part, *filename)
	case 4:
		day04.Run(*part, *filename)
	case 5:
		day05.Run(*part, *filename)
	case 6:
		day06.Run(*part, *filename)
	case 7:
		day07.Run(*part, *filename)
	case 8:
		day08.Run(*part, *filename)
	case 9:
		day09.Run(*part, *filename)
	case 10:
		day10.Run(*part, *filename)
	case 11:
		day11.Run(*part, *filename)
	case 12:
		day12.Run(*part, *filename)
	case 13:
		day13.Run(*part, *filename)
	default:
		fmt.Println("Invalid day. Please select a day between 1 and 25.")
	}
}
