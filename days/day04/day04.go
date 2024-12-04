package day04

import (
	"fmt"
)

func part1(input []string) {
	var count int

	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if y == 0 && x == 0 {
				continue
			}
			count += countXMAS(y, x, input)
		}
	}

	fmt.Println("Result Day 4, Part 1:", count)
}

func countXMAS(dirY int, dirX int, input []string) int {
	var count, startY, endY, startX, endX int

	if dirY == 1 {
		startY = 0
		endY = len(input[0]) - 3
	} else if dirY == -1 {
		startY = 3
		endY = len(input)
	} else {
		startY = 0
		endY = len(input)
	}

	if dirX == 1 {
		startX = 0
		endX = len(input[0]) - 3
	} else if dirX == -1 {
		startX = 3
		endX = len(input[0])
	} else {
		startX = 0
		endX = len(input[0])
	}

	y := startY
	for y < endY {
		x := startX
		for x < endX {
			if input[y][x] == 'X' && input[y+dirY][x+dirX] == 'M' && input[y+dirY*2][x+dirX*2] == 'A' && input[y+dirY*3][x+dirX*3] == 'S' {
				count++
			}
			x += 1
		}
		y += 1
	}

	return count
}

func part2(input []string) {
	var count int

	y := 1
	for y < len(input)-1 {
		x := 1
		for x < len(input[y])-1 {
			if input[y][x] == 'A' && input[y-1][x-1] == 'M' && input[y-1][x+1] == 'M' && input[y+1][x-1] == 'S' && input[y+1][x+1] == 'S' {
				count++
			}
			if input[y][x] == 'A' && input[y-1][x-1] == 'S' && input[y-1][x+1] == 'S' && input[y+1][x-1] == 'M' && input[y+1][x+1] == 'M' {
				count++
			}
			if input[y][x] == 'A' && input[y-1][x-1] == 'M' && input[y-1][x+1] == 'S' && input[y+1][x-1] == 'M' && input[y+1][x+1] == 'S' {
				count++
			}
			if input[y][x] == 'A' && input[y-1][x-1] == 'S' && input[y-1][x+1] == 'M' && input[y+1][x-1] == 'S' && input[y+1][x+1] == 'M' {
				count++
			}
			x += 1
		}
		y += 1
	}

	fmt.Println("Result Day 4, Part 2:", count)
}
