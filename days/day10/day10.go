package day10

import (
	"fmt"
	"time"

	"github.com/mbe81/advent-of-code-2024/lib/convert"
)

func part1(input []string) {
	t := time.Now()

	grid := parseGrid(input)
	maxY, maxX := len(grid), len(grid[0])

	var totalScore int
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 0 {
				trailheadsFound := walk(grid, nil, 0, y, x, maxY, maxX)
				totalScore += len(trailheadsFound)
			}
		}
	}

	fmt.Println("Result day 9, part 1:", totalScore, "- duration:", time.Since(t))
}

func part2(input []string) {
	t := time.Now()

	grid := parseGrid(input)
	maxY, maxX := len(grid), len(grid[0])

	var totalScore int
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 0 {
				trailheadsFound := walk(grid, nil, 0, y, x, maxY, maxX)
				for _, v := range trailheadsFound {
					totalScore += v
				}
			}
		}
	}

	fmt.Println("Result day 9, part 2:", totalScore, "- duration:", time.Since(t))
}

func parseGrid(input []string) [][]int {
	grid := make([][]int, len(input))
	for y := range input {
		grid[y] = make([]int, len(input[y]))
		for x := range input[y] {
			grid[y][x] = convert.StringToInt(input[y][x : x+1])
		}
	}
	return grid
}

type Point struct {
	Y int
	X int
}

var directions = []Point{
	{Y: -1, X: 0},
	{Y: 1, X: 0},
	{Y: 0, X: -1},
	{Y: 0, X: 1},
}

func walk(grid [][]int, trailheadsFound map[Point]int, currentHeight, currentY, currentX, maxY, maxX int) map[Point]int {
	if currentHeight == 9 {
		if trailheadsFound == nil {
			trailheadsFound = make(map[Point]int)
		}
		trailheadsFound[Point{Y: currentY, X: currentX}] += 1
		return trailheadsFound
	}

	nextHeight := currentHeight + 1
	for _, dir := range directions {
		nextY, nextX := currentY+dir.Y, currentX+dir.X
		if withinBoundaries(nextY, nextX, maxY, maxX) && grid[nextY][nextX] == nextHeight {
			trailheadsFound = walk(grid, trailheadsFound, nextHeight, nextY, nextX, maxY, maxX)
		}
	}

	return trailheadsFound
}

func withinBoundaries(y, x, maxY, maxX int) bool {
	return y >= 0 && y < maxY && x >= 0 && x < maxX
}
