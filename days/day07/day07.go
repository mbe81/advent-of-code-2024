package day07

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/mbe81/advent-of-code-2024/lib/convert"
)

func part1(input []string) {
	t := time.Now()

	testValues, numbers := parseInput(input)
	operators := []string{"+", "*"}

	var sumTestValues int
	for i := range testValues {
		lineOperators := generateOperators(operators, len(numbers[i])-1)
		for j := range lineOperators {
			calculatedValue := calculateValue(numbers[i], lineOperators[j], testValues[i])
			if calculatedValue == testValues[i] {
				sumTestValues += calculatedValue
				break
			}
		}
	}

	fmt.Println("Result day 7, part 1:", sumTestValues, "- duration:", time.Since(t))
}

func part2(input []string) {
	t := time.Now()

	testValues, numbers := parseInput(input)
	operators := []string{"+", "*", "|"}

	var sumTestValues int
	for i := range testValues {
		lineOperators := generateOperators(operators, len(numbers[i])-1)
		for j := range lineOperators {
			calculatedValue := calculateValue(numbers[i], lineOperators[j], testValues[i])
			if calculatedValue == testValues[i] {
				sumTestValues += calculatedValue
				break
			}
		}
	}

	fmt.Println("Result day 7, part 2:", sumTestValues, "- duration:", time.Since(t))
}

func parseInput(input []string) ([]int, [][]int) {
	results, numbers := make([]int, len(input)), make([][]int, len(input))
	for y := range input {
		parts := strings.Split(input[y], ":")
		results[y] = convert.StringToInt(parts[0])
		numbers[y] = convert.LineToInts(parts[1], nil)
	}
	return results, numbers
}

func generateOperators(operators []string, length int) []string {
	if length == 0 {
		return []string{""}
	}

	subOperators := generateOperators(operators, length-1)
	var result []string
	for i := range operators {
		for j := range subOperators {
			result = append(result, operators[i]+subOperators[j])
		}
	}
	return result
}

func calculateValue(numbers []int, operators string, testValue int) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if result > testValue {
			return -1
		}
		switch operators[i-1] {
		case '+':
			result += numbers[i]
		case '*':
			result *= numbers[i]
		case '|':
			result = concatenateNumbers(result, numbers[i])
		}
	}
	return result
}

func concatenateNumbers(a, b int) int {
	digits := int(math.Log10(float64(b)) + 1)
	return a*int(math.Pow10(digits)) + b
}
