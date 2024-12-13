package day13

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/mbe81/advent-of-code-2024/lib/convert"
)

func part1(input []string) {
	t := time.Now()

	machines := parseInput(input)

	totalTokens := 0
	for _, machine := range machines {
		minTokens := math.MaxInt64
		for a := 0; a <= 100; a++ {
			for b := 0; b <= 100; b++ {
				tokens := a*3 + b*1
				x := a*machine.ButtonA.X + b*machine.ButtonB.X
				y := a*machine.ButtonA.Y + b*machine.ButtonB.Y
				if x == machine.Prize.X && y == machine.Prize.Y {
					if tokens < minTokens {
						minTokens = tokens
					}
					continue
				}
				if x > machine.Prize.X || y > machine.Prize.Y || tokens > minTokens {
					break
				}
			}
		}
		if minTokens != math.MaxInt64 {
			totalTokens += minTokens
		}
	}

	fmt.Println("Result day 13, part 1:", totalTokens, "- duration:", time.Since(t))
}

func part2(input []string) {
	t := time.Now()

	machines := parseInput(input)

	totalTokens := 0
	for _, machine := range machines {
		machine.Prize.X += 10_000_000_000_000
		machine.Prize.Y += 10_000_000_000_000
		a, b, err := solveEquation(float64(machine.ButtonA.X), float64(machine.ButtonA.Y), float64(machine.ButtonB.X), float64(machine.ButtonB.Y), float64(machine.Prize.X), float64(machine.Prize.Y))
		if err == nil {
			totalTokens += int(a*3 + b*1)
		}
	}

	fmt.Println("Result day 13, part 2:", totalTokens, "- duration:", time.Since(t))
}

type Point struct {
	X int
	Y int
}

type Machine struct {
	ButtonA Point
	ButtonB Point
	Prize   Point
}

func parseInput(input []string) []Machine {
	machines := make([]Machine, 0)
	var machine Machine
	for _, line := range input {
		if line == "" {
			continue
		} else if line[:8] == "Button A" {
			machine.ButtonA.X = convert.StringToInt(line[strings.Index(line, "X")+2 : strings.Index(line, ",")])
			machine.ButtonA.Y = convert.StringToInt(line[strings.Index(line, "Y")+2:])
		} else if line[:8] == "Button B" {
			machine.ButtonB.X = convert.StringToInt(line[strings.Index(line, "X")+2 : strings.Index(line, ",")])
			machine.ButtonB.Y = convert.StringToInt(line[strings.Index(line, "Y")+2:])
		} else if line[:5] == "Prize" {
			machine.Prize.X = convert.StringToInt(line[strings.Index(line, "X")+2 : strings.Index(line, ",")])
			machine.Prize.Y = convert.StringToInt(line[strings.Index(line, "Y")+2:])
			machines = append(machines, machine)
		}
	}
	return machines
}

func solveEquation(xa, ya, xb, yb, xp, yp float64) (float64, float64, error) {
	b := (yp - xp*ya/xa) / (yb - (xb * ya / xa))
	a := (xp - b*xb) / xa

	a, b = math.Round(a), math.Round(b)

	if a*xa+b*xb == xp && a*ya+b*yb == yp && a >= 0 && b >= 0 {
		return a, b, nil
	} else {
		return 0, 0, fmt.Errorf("no solution found")
	}
}
