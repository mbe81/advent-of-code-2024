package day08

import (
	"fmt"
	"slices"
	"time"
)

type Point struct {
	Y int
	X int
}

func part1(input []string) {
	t := time.Now()

	nodes := slices.Clone(input)

	maxY, maxX := len(input), len(input[0])
	nodeY, nodeX := 0, 0

	frequencies := findFrequencies(input)
	for _, frequency := range frequencies {
		locations := findLocations(input, frequency)

		for i := 0; i < len(locations)-1; i++ {
			for j := i + 1; j < len(locations); j++ {
				diffY, diffX := locations[j].Y-locations[i].Y, locations[j].X-locations[i].X

				nodeY, nodeX = locations[i].Y-diffY, locations[i].X-diffX
				if withinBoundaries(nodeY, nodeX, maxY, maxX) {
					nodes[nodeY] = nodes[nodeY][:nodeX] + "#" + nodes[nodeY][nodeX+1:]
				}

				nodeY, nodeX = locations[j].Y+diffY, locations[j].X+diffX
				if withinBoundaries(nodeY, nodeX, maxY, maxX) {
					nodes[nodeY] = nodes[nodeY][:nodeX] + "#" + nodes[nodeY][nodeX+1:]
				}
			}
		}
	}

	fmt.Println("Result day 8, part 1:", countNodes(nodes), "- duration:", time.Since(t))
}

func part2(input []string) {
	t := time.Now()

	nodes := slices.Clone(input)

	maxY, maxX := len(input), len(input[0])
	nodeY, nodeX := 0, 0

	frequencies := findFrequencies(input)
	for _, frequency := range frequencies {
		locations := findLocations(input, frequency)

		for i := 0; i < len(locations)-1; i++ {
			for j := i + 1; j < len(locations); j++ {
				diffY, diffX := locations[j].Y-locations[i].Y, locations[j].X-locations[i].X

				nodeY, nodeX = locations[i].Y, locations[i].X
				for {
					if !withinBoundaries(nodeY, nodeX, maxY, maxX) {
						break
					}
					nodes[nodeY] = nodes[nodeY][:nodeX] + "#" + nodes[nodeY][nodeX+1:]
					nodeY, nodeX = nodeY-diffY, nodeX-diffX
				}

				nodeY, nodeX = locations[i].Y, locations[i].X
				for {
					if !withinBoundaries(nodeY, nodeX, maxY, maxX) {
						break
					}
					nodes[nodeY] = nodes[nodeY][:nodeX] + "#" + nodes[nodeY][nodeX+1:]
					nodeY, nodeX = nodeY+diffY, nodeX+diffX
				}
			}
		}
	}

	fmt.Println("Result day 8, part 2:", countNodes(nodes), "- duration:", time.Since(t))
}

func findFrequencies(input []string) []uint8 {
	var frequencies []uint8
	for y := range input {
		for x := range input[y] {
			if input[y][x] != '.' {
				if !slices.Contains(frequencies, input[y][x]) {
					frequencies = append(frequencies, input[y][x])
				}
			}
		}
	}
	return frequencies
}

func findLocations(input []string, frequency uint8) []Point {
	var locations []Point
	for y := range input {
		for x := range input[y] {
			if input[y][x] == frequency {
				locations = append(locations, Point{Y: y, X: x})
			}
		}
	}
	return locations
}

func countNodes(input []string) int {
	var nodes int
	for y := range input {
		for x := range input[y] {
			if input[y][x] == '#' {
				nodes++
			}
		}
	}
	return nodes
}

func withinBoundaries(y, x, maxY, maxX int) bool {
	return y >= 0 && y < maxY && x >= 0 && x < maxX
}
