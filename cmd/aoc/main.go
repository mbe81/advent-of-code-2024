package main

import (
	"flag"
	"fmt"

	"github.com/mbe81/advent-of-code-2024/days/day01"
	"github.com/mbe81/advent-of-code-2024/days/day02"
	"github.com/mbe81/advent-of-code-2024/days/day03"
	"github.com/mbe81/advent-of-code-2024/days/day04"
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
	}
}
