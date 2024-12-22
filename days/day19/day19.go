package day19

import (
	"fmt"
	"strings"
	"time"
)

func part1(input []string) {
	t := time.Now()

	towels, designs := parseInput(input)
	count := 0
	for _, design := range designs {
		cache = make(map[int]int)
		if findTowel(design, towels, 0) > 0 {
			count++
		}
	}

	fmt.Println("Result day 19, part 1:", count, "- duration:", time.Since(t))
}

func part2(input []string) {
	t := time.Now()

	towels, designs := parseInput(input)
	count := 0
	for _, design := range designs {
		cache = make(map[int]int)
		count += findTowel(design, towels, 0)
	}

	fmt.Println("Result day 19, part 2:", count, "- duration:", time.Since(t))
}

func parseInput(input []string) (towels []string, designs []string) {
	towels = strings.Split(input[0], ", ")
	designs = input[2:]

	return towels, designs
}

var cache map[int]int

func findTowel(design string, availableTowels []string, pos int) int {
	if pos == len(design) {
		return 1
	}
	if _, ok := cache[pos]; ok {
		return cache[pos]
	}

	count := 0
	for _, towel := range availableTowels {
		if strings.HasPrefix(design[pos:], towel) {
			count += findTowel(design, availableTowels, pos+len(towel))
		}
	}
	cache[pos] = count

	return count
}
