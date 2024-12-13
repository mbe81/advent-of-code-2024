package day12

import (
	"fmt"
	"time"
)

func part1(input []string) {
	t := time.Now()

	visited := make(map[Point]bool)
	regions := make(map[RegionID]RegionValue)

	maxY, maxX := len(input), len(input[0])
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if visited[Point{y, x}] {
				continue
			}
			regionID := RegionID{Plant: string(input[y][x]), Y: y, X: x}
			visited, regions = exploreRegion(input, visited, regions, regionID, y, x, maxY, maxX)
		}
	}

	totalPrice := 0
	for _, region := range regions {
		totalPrice += region.Area * region.Perimeter
	}

	fmt.Println("Result day 12, part 1:", totalPrice, "- duration:", time.Since(t))
}

func part2(input []string) {
	t := time.Now()

	visited := make(map[Point]bool)
	regions := make(map[RegionID]RegionValue)

	maxY, maxX := len(input), len(input[0])
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if visited[Point{y, x}] {
				continue
			}
			regionID := RegionID{Plant: string(input[y][x]), Y: y, X: x}
			visited, regions = exploreRegion(input, visited, regions, regionID, y, x, maxY, maxX)
		}
	}

	totalPrice := 0
	for _, region := range regions {
		totalPrice += region.Area * region.Sides
	}

	fmt.Println("Result day 12, part 2:", totalPrice, "- duration:", time.Since(t))
}

type Point struct {
	Y int
	X int
}

type RegionID struct {
	Plant string
	Y     int
	X     int
}

type RegionValue struct {
	Area      int
	Perimeter int
	Sides     int
}

func exploreRegion(input []string, visited map[Point]bool, region map[RegionID]RegionValue, regionID RegionID, y, x, maxY, maxX int) (map[Point]bool, map[RegionID]RegionValue) {
	if !withinBoundaries(y, x, len(input), len(input[0])) {
		return visited, region
	}
	if visited[Point{y, x}] {
		return visited, region
	}
	if string(input[y][x]) != regionID.Plant {
		return visited, region
	}

	visited[Point{y, x}] = true

	region[regionID] = RegionValue{
		Area:      region[regionID].Area + 1,
		Perimeter: region[regionID].Perimeter + countPerimeters(input, y, x, maxY, maxX),
		Sides:     region[regionID].Sides + countCorners(input, y, x, maxY, maxX),
	}

	visited, region = exploreRegion(input, visited, region, regionID, y-1, x, maxY, maxX)
	visited, region = exploreRegion(input, visited, region, regionID, y, x-1, maxY, maxX)
	visited, region = exploreRegion(input, visited, region, regionID, y+1, x, maxY, maxX)
	visited, region = exploreRegion(input, visited, region, regionID, y, x+1, maxY, maxX)

	return visited, region
}

var perimeterDirections = []Point{{Y: -1, X: 0}, {Y: 1, X: 0}, {Y: 0, X: -1}, {Y: 0, X: 1}}

func countPerimeters(input []string, y, x, maxY, maxX int) int {
	perimeters := 0
	for _, dir := range perimeterDirections {
		if !withinBoundaries(y+dir.Y, x+dir.X, maxY, maxX) || input[y+dir.Y][x+dir.X] != input[y][x] {
			perimeters++
		}
	}
	return perimeters
}

var sideDirections = []Point{{Y: -1, X: -1}, {Y: -1, X: 1}, {Y: 1, X: -1}, {Y: 1, X: 1}}

func countCorners(input []string, y, x, maxY, maxX int) int {
	sides := 0

	for _, dir := range sideDirections {
		if isCorner(input, y, x, dir.Y, dir.X, maxY, maxX) {
			sides++
		}
	}

	return sides
}

func isCorner(input []string, y, x, dirY, dirX, maxY, maxX int) bool {
	if (!withinBoundaries(y+dirY, x, maxY, maxX) || input[y+dirY][x] != input[y][x]) &&
		(!withinBoundaries(y, x+dirX, maxY, maxX) || input[y][x+dirX] != input[y][x]) {
		return true // Outer corner
	}
	if (withinBoundaries(y+dirY, x, maxY, maxX) && input[y+dirY][x] == input[y][x]) &&
		(withinBoundaries(y, x+dirX, maxY, maxX) && input[y][x+dirX] == input[y][x]) &&
		(withinBoundaries(y+dirY, x+dirX, maxY, maxX) && input[y+dirY][x+dirX] != input[y][x]) {
		return true // Inner corner
	}
	return false
}

func withinBoundaries(y, x, maxY, maxX int) bool {
	return y >= 0 && y < maxY && x >= 0 && x < maxX
}
