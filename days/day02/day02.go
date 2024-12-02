package day02

import (
	"fmt"

	"github.com/mbe81/advent-of-code-2024/lib/calc"
	"github.com/mbe81/advent-of-code-2024/lib/convert"
)

func part1(input []string) {
	var safeReports int
	for i := range input {
		levels := convert.LineToInts(input[i])
		if validateReport(levels) {
			safeReports++
		}
	}

	fmt.Println("Result Day 2, Part 1:", safeReports)
}

func part2(input []string) {
	var safeReports int
	for i := range input {
		levels := convert.LineToInts(input[i])
		if validateReport(levels) {
			safeReports++
			continue
		}
		for j := 0; j < len(levels); j++ {
			adjustedLevels := make([]int, 0, len(levels)-1)
			adjustedLevels = append(adjustedLevels, levels[:j]...)
			adjustedLevels = append(adjustedLevels, levels[j+1:]...)

			if validateReport(adjustedLevels) {
				safeReports++
				break
			}
		}
	}

	fmt.Println("Result Day 2, Part 2:", safeReports)
}

func validateReport(levels []int) bool {
	safeReport := true
	expectedSign := calc.Sign(levels[len(levels)-1] - levels[0])

	for j := 0; j < len(levels)-1; j++ {
		sign := calc.Sign(levels[j+1] - levels[j])
		if sign != expectedSign {
			safeReport = false
			break
		}
		diff := calc.AbsDifference(levels[j+1], levels[j])
		if diff < 1 || diff > 3 {
			safeReport = false
			break
		}
	}
	return safeReport
}
