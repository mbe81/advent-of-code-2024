package day03

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mbe81/advent-of-code-2024/lib/convert"
)

func part1(input []string) {
	var instructionPattern = regexp.MustCompile(`mul\(\d+,\d+\)`)

	var result int
	for _, command := range input {
		instructions := instructionPattern.FindAllString(command, -1)

		for _, instruction := range instructions {
			result += multiplyInstr(instruction)
		}
	}

	fmt.Println("Result Day 3, Part 1:", result)
}

func part2(input []string) {
	var instructionPattern = regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)

	var result int
	var enabled = true
	for _, command := range input {
		instructions := instructionPattern.FindAllString(command, -1)

		for _, instruction := range instructions {
			if instruction == "do()" {
				enabled = true
				continue
			}
			if instruction == "don't()" {
				enabled = false
				continue
			}
			if enabled {
				result += multiplyInstr(instruction)
			}
		}
	}

	fmt.Println("Result Day 3, Part 2:", result)
}

func multiplyInstr(instruction string) int {
	multipliersRaw := strings.Replace(instruction[4:len(instruction)-1], ",", " ", 1)
	multipliers := convert.LineToInts(multipliersRaw)
	return multipliers[0] * multipliers[1]
}
