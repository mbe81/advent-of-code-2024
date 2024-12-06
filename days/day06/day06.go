package day06

import (
	"fmt"
)

type Visit struct {
	Position  Point
	Direction Point
}

type Point struct {
	Y int
	X int
}

var (
	Left  = Point{Y: 0, X: -1}
	Right = Point{Y: 0, X: 1}
	Up    = Point{Y: -1, X: 0}
	Down  = Point{Y: 1, X: 0}
)

func part1(input []string) {
	start := getStartingPosition(input)
	visited, _ := solve(input, start, nil)

	fmt.Println("Result Day 6, Part 1:", countUniqueVisits(visited))
}

func part2(input []string) {
	var loopCounter int
	start := getStartingPosition(input)

	for y := range input {
		for x := range input[y] {
			if input[y][x] == '#' || input[y][x] == '^' {
				continue
			}

			obstruction := Point{Y: y, X: x}

			_, loop := solve(input, start, &obstruction)
			if loop {
				loopCounter++
			}
		}
	}

	fmt.Println("Result Day 6, Part 2:", loopCounter)
}

func getStartingPosition(input []string) Point {
	for y := range input {
		for x := range input[y] {
			if input[y][x] == '^' {
				return Point{Y: y, X: x}
			}
		}
	}
	return Point{}
}

func solve(input []string, start Point, obstruction *Point) (map[Visit]bool, bool) {
	pos := start
	dir := Up

	maxY := len(input) - 1
	maxX := len(input[0]) - 1

	visited := make(map[Visit]bool)
	visited[Visit{pos, dir}] = true

	loopDetected := false
	for {
		if pos.Y+dir.Y < 0 || pos.Y+dir.Y > maxY || pos.X+dir.X < 0 || pos.X+dir.X > maxX {
			break
		}
		for {
			if input[pos.Y+dir.Y][pos.X+dir.X] == '#' || (obstruction != nil && *obstruction == Point{pos.Y + dir.Y, pos.X + dir.X}) {
				dir = turnRight(dir)
			} else {
				break
			}
		}
		pos.Y, pos.X = pos.Y+dir.Y, pos.X+dir.X
		if visited[Visit{pos, dir}] {
			loopDetected = true
			break
		}
		visited[Visit{pos, dir}] = true
	}
	return visited, loopDetected
}

func turnRight(dir Point) Point {
	switch dir {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	}
	return Point{}
}

func countUniqueVisits(visited map[Visit]bool) int {
	visitedUnique := make(map[Point]bool)
	for v := range visited {
		visitedUnique[Point{Y: v.Position.Y, X: v.Position.X}] = true
	}
	return len(visitedUnique)
}
