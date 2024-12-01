package file

import (
	"bufio"
	"os"
)

// ReadLines reads a file and returns its lines as a slice of strings. Panics if the file cannot be read.
func ReadLines(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
