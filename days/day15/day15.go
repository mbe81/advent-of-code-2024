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
			for x := range input[y] {
				if input[y][x] == '#' || input[y][x] == '.' {
					grid[y][x*2], grid[y][x*2+1] = input[y][x], input[y][x]
				} else if input[y][x] == 'O' {
					grid[y][x*2], grid[y][x*2+1] = '[', ']'
				} else if input[y][x] == '@' {
					grid[y][x*2], grid[y][x*2+1] = '@', '.'
				}
			}
		} else {
			grid = append(grid, make([]uint8, len(input[y])))
			for x := range input[y] {
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

func part2(input []string) {
	t := time.Now()

	grid, moves := parseInput(input, true)

	x, y := findRobot(grid)
	for _, move := range moves {
		grid, x, y = moveRobot(grid, x, y, uint8(move))
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

func moveRobot(grid [][]uint8, x, y int, move uint8) ([][]uint8, int, int) {
	if move == '<' || move == '>' {
		return moveHorizontal(grid, x, y, move)
	} else if move == '^' || move == 'v' {
		return moveVertical(grid, x, y, move)
	}
	return grid, x, y
}

func moveHorizontal(grid [][]uint8, x, y int, move uint8) ([][]uint8, int, int) {
	var dirX int
	if move == '<' {
		dirX = -1
	} else if move == '>' {
		dirX = 1
	}

	checkX := x
	for {
		checkX = checkX + dirX
		if grid[y][checkX] == '#' {
			return grid, x, y
		}
		if grid[y][checkX] == '.' {
			break
		}
	}

	for checkX != x {
		grid[y][checkX] = grid[y][checkX-dirX]
		checkX = checkX - dirX
	}
	grid[y][x] = '.'

	return grid, x + dirX, y
}

func moveVertical(grid [][]uint8, x, y int, move uint8) ([][]uint8, int, int) {
	var dirY int
	if move == '^' {
		dirY = -1
	} else if move == 'v' {
		dirY = 1
	}

	if grid[y+dirY][x] == '#' {
		return grid, x, y
	}

	if grid[y+dirY][x] == '.' {
		grid[y+dirY][x] = grid[y][x]
		grid[y][x] = '.'
		return grid, x, y + dirY
	}

	boxes = make(map[block]bool)

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
		minY = min(minY, box.y)
		maxY = max(maxY, box.y)
	}

	var checkY, endY int
	if dirY == -1 {
		checkY, endY = minY, maxY
	} else if dirY == 1 {
		checkY, endY = maxY, minY
	}

	for checkY != endY-dirY {
		for box := range boxes {
			if box.y == checkY {
				grid[checkY+dirY][box.x] = grid[checkY][box.x]
				grid[checkY][box.x] = '.'
			}
		}
		checkY -= dirY
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

	if grid[y+dirY][x] == 'O' {
		allow := allowMove(grid, x, y+dirY, dirY)
		if allow {
			boxes[block{x, y}] = true
		}
		return allow
	}

	if grid[y+dirY][x] == '[' {
		left := allowMove(grid, x, y+dirY, dirY)
		right := allowMove(grid, x+1, y+dirY, dirY)
		if left && right {
			boxes[block{x, y}] = true
		}
		return left && right
	}
	if grid[y+dirY][x] == ']' {
		left := allowMove(grid, x-1, y+dirY, dirY)
		right := allowMove(grid, x, y+dirY, dirY)
		if left && right {
			boxes[block{x, y}] = true
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
