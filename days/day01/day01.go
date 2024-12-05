package day01

import (
	"fmt"
	"sort"

	"github.com/mbe81/advent-of-code-2024/lib/calc"
	"github.com/mbe81/advent-of-code-2024/lib/convert"
)

func part1(input []string) {
	left, right := make([]int, len(input)), make([]int, len(input))
	for i := range input {
		numbers := convert.LineToInts(input[i], nil)
		left[i] = numbers[0]
		right[i] = numbers[1]
	}

	sort.Ints(left)
	sort.Ints(right)

	var distance int
	for i := range left {
		distance += calc.AbsDifference(left[i], right[i])
	}

	fmt.Println("Result Day 1, Part 1:", distance)
}

func part2(input []string) {
	left, right := make([]int, len(input)), make([]int, len(input))
	for i := range input {
		numbers := convert.LineToInts(input[i], nil)
		left[i] = numbers[0]
		right[i] = numbers[1]
	}

	var similarity int
	for i := range left {
		for j := range right {
			if left[i] == right[j] {
				similarity += left[i]
			}
		}
	}

	fmt.Println("Result Day 1, Part 2:", similarity)
}
