package day18

import (
	"fmt"
	"time"

	"github.com/mbe81/advent-of-code-2024/lib/convert"
)

var separator = ","
var gridSize = 71

func part1(input []string) {
	t := time.Now()

	start := Point{0, 0}
	end := Point{gridSize - 1, gridSize - 1}

	grid := getGrid(gridSize, gridSize)

	for i, line := range input {
		if i >= 1024 {
			break
		}
		n := convert.LineToInts(line, &separator)
		x, y := n[0], n[1]
		grid[x][y] = '#'
	}

	steps := len(bfsShortestPath(grid, start, end)) - 1

	fmt.Println("Result day 18, part 1:", steps, "- duration:", time.Since(t))
}

func part2(input []string) {
	t := time.Now()

	start := Point{0, 0}
	end := Point{gridSize - 1, gridSize - 1}

	lowerBound := 1024
	upperbound := len(input)
	var blocks = (lowerBound + upperbound) / 2

	for lowerBound < upperbound {

		grid := getGrid(gridSize, gridSize)
		for i, line := range input {
			if i >= blocks {
				break
			}
			n := convert.LineToInts(line, &separator)
			x, y := n[0], n[1]
			grid[x][y] = '#'
		}
		path := bfsShortestPath(grid, start, end)
		if path == nil {
			upperbound = blocks
		} else {
			lowerBound = blocks + 1
		}
		blocks = (lowerBound + upperbound) / 2
	}

	fmt.Println("Result day 18, part 2:", input[blocks-1], "- duration:", time.Since(t))
}

func getGrid(sizeX, sizeY int) [][]uint8 {
	var grid = make([][]uint8, sizeY)
	for x := range grid {
		grid[x] = make([]uint8, sizeX)
		for y := range grid[x] {
			grid[x][y] = '.'
		}
	}

	return grid
}

type Point struct {
	X, Y int
}

type Move struct {
	Pos Point
	Dir Point
}

func bfsShortestPath(grid [][]uint8, start, end Point) []Point {
	dirs := []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	rows, cols := len(grid), len(grid[0])

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	parents := make(map[Point]Point)
	queue := []Point{start}
	visited[start.X][start.Y] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur == end {
			path := make([]Point, 0)
			for p := end; p != start; p = parents[p] {
				path = append([]Point{p}, path...)
			}
			return append([]Point{start}, path...)
		}

		for _, d := range dirs {
			r, c := cur.X+d.X, cur.Y+d.Y
			if !withinGrid(r, c) {
				continue
			}
			if grid[r][c] == '#' {
				continue
			}
			if visited[r][c] {
				continue
			}
			visited[r][c] = true
			parents[Point{r, c}] = cur
			queue = append(queue, Point{r, c})
		}
	}
	return nil
}

func withinGrid(y, x int) bool {
	return y >= 0 && y < gridSize && x >= 0 && x < gridSize
}
