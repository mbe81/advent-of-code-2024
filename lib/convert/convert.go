package convert

import (
	"strconv"
	"strings"
)

// StringToInt converts a string to an integer. Panics if the conversion fails.
func StringToInt(s string) int {
	i, err := strconv.Atoi(strings.Trim(s, " "))
	if err != nil {
		panic(err)
	}
	return i
}

// LineToInts converts a line to multiple numbers. Panics if the conversion fails.
func LineToInts(line string) []int {
	var numbers []int
	for _, n := range strings.Fields(line) {
		numbers = append(numbers, StringToInt(n))
	}
	return numbers
}
