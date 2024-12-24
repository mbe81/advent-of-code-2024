package day24

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/mbe81/advent-of-code-2024/lib/convert"
)

func part1(input []string) {
	t := time.Now()

	wires, ports := parseInput(input)
	wires, ports = powerCircuit(wires, ports)
	_, _, result := getResults(wires)

	fmt.Println("Result day 24, part 1:", result, "- duration:", time.Since(t))
}

func part2(input []string) {
	t := time.Now()

	// TODO: No part 2 for day 24 until now :-(

	fmt.Println("Result day 24, part 1:", 0, "- duration:", time.Since(t))
}

type Port struct {
	In1       string
	In2       string
	Out       string
	Operator  string
	Connected bool
}

func parseInput(input []string) (map[string]int, []Port) {
	wires := make(map[string]int)
	ports := make([]Port, 0)

	for _, line := range input {
		if strings.Index(line, ":") > 0 {
			split := strings.Split(line, ": ")
			wires[split[0]] = convert.StringToInt(split[1])
		} else if strings.Index(line, "->") > 0 {
			split := strings.Split(line, " ")
			ports = append(ports, Port{
				In1:      split[0],
				In2:      split[2],
				Out:      split[4],
				Operator: split[1],
			})
		}
	}

	return wires, ports
}

func powerCircuit(wires map[string]int, ports []Port) (map[string]int, []Port) {
	var newSignals = true
	for newSignals {
		newSignals = false
		for i, _ := range ports {
			if ports[i].Connected {
				continue
			}
			wire1, ok := wires[ports[i].In1]
			if !ok {
				continue
			}
			wire2, ok := wires[ports[i].In2]
			if !ok {
				continue
			}

			switch ports[i].Operator {
			case "AND":
				wires[ports[i].Out] = wire1 & wire2
			case "OR":
				wires[ports[i].Out] = wire1 | wire2
			case "XOR":
				wires[ports[i].Out] = wire1 ^ wire2
			}

			ports[i].Connected, newSignals = true, true
		}
	}

	return wires, ports
}

func getResults(wires map[string]int) (int, int, int) {
	var inX, inY, outZ int
	for k, v := range wires {
		if k[0] == 'x' {
			bit := convert.StringToInt(k[1:])
			inX += v * int(math.Pow(2, float64(bit)))
		}
		if k[0] == 'y' {
			bit := convert.StringToInt(k[1:])
			inY += v * int(math.Pow(2, float64(bit)))
		}
		if k[0] == 'z' {
			bit := convert.StringToInt(k[1:])
			outZ += v * int(math.Pow(2, float64(bit)))
		}
	}
	return inX, inY, outZ
}
