package day15

import (
	"fmt"
	"math"
	"time"
)

func part1(input []string) {
	t := time.Now()

	grid, moves := parseInput(input, false)

	x, y := findRobot(grid)
	for _, move := range moves {
		grid, x, y = moveRobot(grid, x, y, uint8(move))
	}

	totalSum := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 'O' {
				totalSum += x + y*100
			}
		}
	}

	fmt.Println("Result day 15, part 1:", totalSum, "- duration:", time.Since(t))
}

func parseInput(input []string, exploded bool) ([][]uint8, string) {
	grid := make([][]uint8, 0)
	for y, _ := range input {
		if input[y] == "" {
			break
		}

		if exploded {
			grid = append(grid, make([]uint8, len(input[y])*2))
		} else {
			grid = append(grid, make([]uint8, len(input[y])))
		}

		for x := range input[y] {
			if exploded {
				if input[y][x] == '#' || input[y][x] == '.' {
					grid[y][x*2], grid[y][x*2+1] = input[y][x], input[y][x]
				} else if input[y][x] == 'O' {
					grid[y][x*2], grid[y][x*2+1] = '[', ']'
				} else if input[y][x] == '@' {
					grid[y][x*2], grid[y][x*2+1] = '@', '.'
				}
			} else {
				grid[y][x] = input[y][x]
			}
		}
	}
	var moves string
	for y := len(grid) + 1; y < len(input); y++ {
		moves += input[y]
	}
	return grid, moves
}

func findRobot(grid [][]uint8) (int, int) {
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '@' {
				return x, y
			}
		}
	}
	return -1, -1
}

func moveRobot(grid [][]uint8, x, y int, move uint8) ([][]uint8, int, int) {
	var dirX, dirY int
	switch move {
	case '^':
		dirX, dirY = 0, -1
	case 'v':
		dirX, dirY = 0, 1
	case '<':
		dirX, dirY = -1, 0
	case '>':
		dirX, dirY = 1, 0
	}

	checkX, checkY := x, y
	for {
		checkX, checkY = checkX+dirX, checkY+dirY
		if grid[checkY][checkX] == '#' {
			return grid, x, y
		}
		if grid[checkY][checkX] == '.' {
			break
		}
	}

	for checkX != x || checkY != y {
		grid[checkY][checkX] = grid[checkY-dirY][checkX-dirX]
		checkX, checkY = checkX-dirX, checkY-dirY
	}
	grid[y][x] = '.'

	return grid, x + dirX, y + dirY
}

func part2(input []string) {
	t := time.Now()

	grid, moves := parseInput(input, true)

	x, y := findRobot(grid)
	for _, move := range moves {
		grid, x, y = moveRobotPart2(grid, x, y, uint8(move))
	}

	totalSum := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '[' {
				totalSum += x + y*100
			}
		}
	}

	fmt.Println("Result day 15, part 2:", totalSum, "- duration:", time.Since(t))
}

type block struct {
	x, y int
}

var boxes map[block]bool

func moveRobotPart2(grid [][]uint8, x, y int, move uint8) ([][]uint8, int, int) {
	if move == '<' || move == '>' {
		return moveRobot(grid, x, y, move)
	}
	var dirY int
	if move == '^' {
		dirY = -1
	} else {
		dirY = 1
	}

	if grid[y+dirY][x] == '#' {
		return grid, x, y
	}
	if grid[y+dirY][x] == '.' {
		return moveRobot(grid, x, y, move)
	}

	boxes = make(map[block]bool, 0)

	if !allowMove(grid, x, y+dirY, dirY) {
		return grid, x, y
	}
	if grid[y+dirY][x] == '[' && !allowMove(grid, x+1, y+dirY, dirY) {
		return grid, x, y
	}
	if grid[y+dirY][x] == ']' && !allowMove(grid, x-1, y+dirY, dirY) {
		return grid, x, y
	}

	var minY, maxY = math.MaxInt, 0
	for box := range boxes {
		if box.y < minY {
			minY = box.y
		}
		if box.y > maxY {
			maxY = box.y
		}
	}

	if dirY == -1 {
		for boxY := minY; boxY <= maxY; boxY++ {
			for box := range boxes {
				if box.y == boxY {
					grid[boxY+dirY][box.x] = grid[boxY][box.x]
					grid[boxY][box.x] = '.'
				}
			}
		}
	}

	if dirY == 1 {
		for boxY := maxY; boxY >= minY; boxY-- {
			for box := range boxes {
				if box.y == boxY {
					grid[boxY+dirY][box.x] = grid[boxY][box.x]
					grid[boxY][box.x] = '.'
				}
			}
		}
	}

	grid[y+dirY][x] = '@'
	grid[y][x] = '.'

	return grid, x, y + dirY
}

func allowMove(grid [][]uint8, x, y, dirY int) bool {

	if grid[y+dirY][x] == '#' {
		return false
	}

	if grid[y+dirY][x] == '.' {
		boxes[block{x, y}] = true
		return true
	}

	if grid[y+dirY][x] == '[' {
		var left, right bool
		left = allowMove(grid, x, y+dirY, dirY)
		right = allowMove(grid, x+1, y+dirY, dirY)

		if left && right {
			if grid[y][x] != ',' {
				boxes[block{x, y}] = true
			}
		}
		return left && right
	}

	if grid[y+dirY][x] == ']' {
		var left, right bool
		left = allowMove(grid, x-1, y+dirY, dirY)
		right = allowMove(grid, x, y+dirY, dirY)

		if left && right {
			if grid[y][x-1] != ',' {
				boxes[block{x, y}] = true
			}
		}
		return left && right
	}

	return false
}

func printGrid(grid [][]uint8) {
	for y := range grid {
		for x := range grid[y] {
			fmt.Print(string(grid[y][x]))
		}
		fmt.Println()
	}
}
