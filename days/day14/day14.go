package day14

import (
	"fmt"
	"strings"
	"time"

	"github.com/mbe81/advent-of-code-2024/lib/convert"
)

func part1(input []string) {
	t := time.Now()

	robots := parseInput(input)

	var maxX, maxY int
	if len(robots) < 100 {
		maxX, maxY = 11, 7
	} else {
		maxX, maxY = 101, 103
	}

	var seconds = 100
	var q1, q2, q3, q4 int
	for i := range robots {
		robots[i].Position.X = (robots[i].Position.X + robots[i].Velocity.X*seconds + maxX*seconds) % maxX
		robots[i].Position.Y = (robots[i].Position.Y + robots[i].Velocity.Y*seconds + maxY*seconds) % maxY

		if robots[i].Position.X < maxX/2 && robots[i].Position.Y < maxY/2 {
			q1++
		} else if robots[i].Position.X > maxX/2 && robots[i].Position.Y < maxY/2 {
			q2++
		} else if robots[i].Position.X < maxX/2 && robots[i].Position.Y > maxY/2 {
			q3++
		} else if robots[i].Position.X > maxX/2 && robots[i].Position.Y > maxY/2 {
			q4++
		}
	}

	safetyFactor := q1 * q2 * q3 * q4

	fmt.Println("Result day 14, part 1:", safetyFactor, "- duration:", time.Since(t))
}

func part2(input []string) {
	t := time.Now()

	robots := parseInput(input)

	var maxX, maxY int
	if len(robots) < 100 {
		maxX, maxY = 11, 7
	} else {
		maxX, maxY = 101, 103
	}

	seconds := 0
	for {
		seconds++
		for i := range robots {
			robots[i].Position.X = (robots[i].Position.X + robots[i].Velocity.X + maxX) % maxX
			robots[i].Position.Y = (robots[i].Position.Y + robots[i].Velocity.Y + maxY) % maxY
		}
		if countOverlap(robots) == 0 {
			break
		}
	}

	fmt.Println("Result day 14, part 2:", seconds, "- duration:", time.Since(t))
}

type Point struct {
	X int
	Y int
}

type Robot struct {
	Position Point
	Velocity Point
}

func parseInput(input []string) []Robot {
	robots := make([]Robot, 0)
	var robot Robot
	for _, line := range input {
		parts := strings.Fields(line)
		pValues := strings.Split(parts[0][2:], ",") // ["0", "4"]
		vValues := strings.Split(parts[1][2:], ",") // ["3", "-3"]

		robot.Position.X = convert.StringToInt(pValues[0])
		robot.Position.Y = convert.StringToInt(pValues[1])
		robot.Velocity.X = convert.StringToInt(vValues[0])
		robot.Velocity.Y = convert.StringToInt(vValues[1])

		robots = append(robots, robot)
	}
	return robots
}

func countOverlap(robots []Robot) int {
	overlap := 0
	for i := range robots {
		for j := i + 1; j < len(robots); j++ {
			if robots[i].Position.X == robots[j].Position.X && robots[i].Position.Y == robots[j].Position.Y {
				overlap++
			}
		}
	}
	return overlap
}

func printRobots(robots []Robot, maxX, maxY int) {
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			found := false
			for _, robot := range robots {
				if robot.Position.X == x && robot.Position.Y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
