package day06

import (
	"fmt"
	"time"
)

const (
	LEFT  int8 = 1
	RIGHT int8 = 2
	UP    int8 = 4
	DOWN  int8 = 8
)

func part1(input []string) {
	t := time.Now()

	startY, startX := getStartingPosition(input)
	visited, _ := solve(input, startY, startX, -1, -1)

	fmt.Println("Result Day 6, Part 1:", countVisits(visited), "Duration:", time.Since(t))
}

func part2(input []string) {
	t := time.Now()

	startY, startX := getStartingPosition(input)
	var loopCounter int
	for y := range input {
		for x := range input[y] {
			if input[y][x] == '#' || input[y][x] == '^' {
				continue
			}
			_, loop := solve(input, startY, startX, int16(y), int16(x))
			if loop {
				loopCounter++
			}
		}
	}

	fmt.Println("Result Day 6, Part 2:", loopCounter, "Duration:", time.Since(t))
}

func getStartingPosition(input []string) (int16, int16) {
	for y := range input {
		for x := range input[y] {
			if input[y][x] == '^' {
				return int16(y), int16(x)
			}
		}
	}
	return 0, 0
}

func solve(input []string, startY, startX, obstrY, obstrX int16) ([][]int8, bool) {
	var posY, posX, newY, newX, maxY, maxX int16
	var dir int8

	posY = startY
	posX = startX
	maxY = int16(len(input) - 1)
	maxX = int16(len(input[0]) - 1)

	dir = UP

	visited := make([][]int8, len(input))
	for i := range visited {
		visited[i] = make([]int8, len(input[i]))
	}

	loopDetected := false
	for {
		newY, newX = move(posY, posX, dir)
		if newY < 0 || newY > maxY || newX < 0 || newX > maxX {
			break
		}
		for {
			if input[newY][newX] == '#' || (newY == obstrY && newX == obstrX) {
				dir = turnRight(dir)
				newY, newX = move(posY, posX, dir)
			} else {
				break
			}
		}
		posY, posX = newY, newX
		if visited[posY][posX]&dir == dir {
			loopDetected = true
			break
		}
		visited[posY][posX] = visited[posY][posX] | dir
	}
	return visited, loopDetected
}

func turnRight(dir int8) int8 {
	switch dir {
	case LEFT:
		return UP
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case UP:
		return RIGHT
	}
	return 0
}

func move(posY, posX int16, dir int8) (int16, int16) {
	switch dir {
	case LEFT:
		return posY, posX - 1
	case RIGHT:
		return posY, posX + 1
	case UP:
		return posY - 1, posX
	case DOWN:
		return posY + 1, posX
	}
	return posY, posX
}

func countVisits(visited [][]int8) int {
	var counter int
	for y := range visited {
		for x := range visited[y] {
			if visited[y][x] > 0 {
				counter++
			}
		}
	}
	return counter
}
