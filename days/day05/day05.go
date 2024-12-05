package day05

import (
	"fmt"
	"strings"

	"github.com/mbe81/advent-of-code-2024/lib/convert"
)

func part1(input []string) {
	var result int

	rules, updates := parseInput(input)

	for _, update := range updates {
		updateOK := true
		for i := range update {
			for j := range update[:i] {
				for _, rule := range rules[update[i]] {
					if rule == update[j] {
						update[j], update[i] = update[i], update[j]
						updateOK = false
					}

				}
			}
		}
		if updateOK {
			result += update[(len(update)-1)/2]
		}
	}

	fmt.Println("Result Day 5, Part 1:", result)
}

func part2(input []string) {
	var result int

	rules, updates := parseInput(input)

	for _, update := range updates {
		updateCorrected := false
		for i := range update {
			for j := range update[:i] {
				for _, rule := range rules[update[i]] {
					if rule == update[j] {
						updateCorrected = true
						update[j], update[i] = update[i], update[j]
					}

				}
			}
		}
		if updateCorrected {
			result += update[(len(update)-1)/2]
		}
	}

	fmt.Println("Result Day 5, Part 2:", result)
}

func parseInput(input []string) (rules map[int][]int, updates [][]int) {
	rules = make(map[int][]int)
	updates = make([][]int, 0)

	for _, line := range input {
		if strings.Contains(line, "|") {
			separator := "|"
			rule := convert.LineToInts(line, &separator)
			rules[rule[0]] = append(rules[rule[0]], rule[1])
		}
		if strings.Contains(line, ",") {
			separator := ","
			update := convert.LineToInts(line, &separator)
			updates = append(updates, update)
		}
	}

	return rules, updates
}
