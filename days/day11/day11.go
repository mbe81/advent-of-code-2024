package day11

import (
	"fmt"
	"math"
	"time"

	"github.com/mbe81/advent-of-code-2024/lib/convert"
)

var separator = " "

func part1(input []string) {
	t := time.Now()

	stones := convert.LineToInts(input[0], &separator)
	totalStones := 0
	for i := 0; i < len(stones); i++ {
		totalStones += countBlinks(stones[i], 25)
	}

	fmt.Println("Result day 11, part 1:", totalStones, "- duration:", time.Since(t))
}

func part2(input []string) {
	t := time.Now()

	stones := convert.LineToInts(input[0], &separator)
	totalStones := 0
	for i := 0; i < len(stones); i++ {
		totalStones += countBlinks(stones[i], 75)
	}

	fmt.Println("Result day 11, part 2:", totalStones, "- duration:", time.Since(t))
}

func splitNumber(n int) (int, int) {
	halfDigits := countDigits(n) / 2
	divisor := int(math.Pow(10, float64(halfDigits)))
	left := n / divisor
	right := n % divisor

	return left, right
}

func countDigits(n int) int {
	if n == 0 {
		return 1
	}
	count := int(math.Log10(float64(n))) + 1
	return count
}

type BlinkState struct {
	stoneValue     int
	remainingSteps int
}

var blinkCache = make(map[BlinkState]int)

func countBlinks(stoneValue, remainingBlinks int) int {
	var count int
	if remainingBlinks == 0 {
		return 1
	}
	if stoneValue == 0 {
		return countBlinks(1, remainingBlinks-1)
	}

	if value, ok := blinkCache[BlinkState{stoneValue, remainingBlinks}]; ok {
		return value
	}

	if countDigits(stoneValue)%2 == 0 {
		left, right := splitNumber(stoneValue)
		count = countBlinks(left, remainingBlinks-1) + countBlinks(right, remainingBlinks-1)
	} else {
		count = countBlinks(stoneValue*2024, remainingBlinks-1)
	}

	if stoneValue < 1000 {
		blinkCache[BlinkState{stoneValue, remainingBlinks}] = count
	}
	
	return count
}
