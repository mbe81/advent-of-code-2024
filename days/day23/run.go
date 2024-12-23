package day23

import (
	"path"
	"runtime"

	"github.com/mbe81/advent-of-code-2024/lib/file"
)

var folder string

func init() {
	_, filename, _, _ := runtime.Caller(0)
	folder = path.Dir(filename)
}

func Run(part int, filename string) {
	var input []string
	input = file.ReadLines(path.Join(folder, filename))

	switch part {
	case 1:
		part1(input)
	case 2:
		part2(input)
	default:
		part1(input)
		part2(input)
	}
}
