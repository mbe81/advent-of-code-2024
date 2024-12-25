package day17

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/mbe81/advent-of-code-2024/lib/convert"
)

var separator = ","

func part1(input []string) {
	t := time.Now()

	computer, instructions, _ := parseInput(input)
	for computer.InstructionPointer < len(instructions) {
		computer = executeInstruction(computer, instructions[computer.InstructionPointer])
	}

	fmt.Println("Result day 17, part 1:", "'"+computer.Output+"'", "- duration:", time.Since(t))
}

func part2(input []string) {
	t := time.Now()

	_, instructions, expected := parseInput(input)
	a := findSolution(instructions, expected, 0, 0)

	fmt.Println("Result day 17, part 2:", a, "- duration:", time.Since(t))
}

type Computer struct {
	RegisterA          int
	RegisterB          int
	RegisterC          int
	InstructionPointer int
	InstructionCount   int
	Output             string
}

type Instruction struct {
	Operator int
	Operand  int
}

func parseInput(input []string) (Computer, []Instruction, []string) {

	var c Computer
	for _, line := range input {
		if line == "" {
			break
		}
		if line[:10] == "Register A" {
			c.RegisterA = convert.StringToInt(line[12:])
		}
		if line[:10] == "Register B" {
			c.RegisterB = convert.StringToInt(line[12:])
		}
		if line[:10] == "Register C" {
			c.RegisterC = convert.StringToInt(line[12:])
		}
	}

	instructions := make([]Instruction, 0)
	split := strings.Split(input[len(input)-1], " ")
	numbers := convert.LineToInts(split[1], &separator)
	for j := 0; j < len(numbers); j = j + 2 {
		instructions = append(instructions,
			Instruction{
				Operator: numbers[j],
				Operand:  numbers[j+1],
			})
	}

	expected := strings.Split(input[len(input)-1][9:], ",")

	return c, instructions, expected
}

func executeInstruction(c Computer, i Instruction) Computer {
	switch i.Operator {
	case 0: // ADV
		c.RegisterA = c.RegisterA / power(2, combo(i.Operand, c))
		c.InstructionPointer++
	case 1: // BXL
		c.RegisterB = c.RegisterB ^ literal(i.Operand)
		c.InstructionPointer++
	case 2: // BST
		c.RegisterB = combo(i.Operand, c) % 8
		c.InstructionPointer++
	case 3: // JNZ
		if c.RegisterA != 0 {
			c.InstructionPointer = literal(i.Operand) / 2
		} else {
			c.InstructionPointer++
		}
	case 4: // BXC
		c.RegisterB = c.RegisterB ^ c.RegisterC
		c.InstructionPointer++
	case 5: // OUT
		if len(c.Output) != 0 {
			c.Output = c.Output + "," + strconv.Itoa(combo(i.Operand, c)%8)
		} else {
			c.Output = strconv.Itoa(combo(i.Operand, c) % 8)
		}
		c.InstructionPointer++
	case 6: // BDV
		c.RegisterB = c.RegisterA / power(2, combo(i.Operand, c))
		c.InstructionPointer++
	case 7: // CDV
		c.RegisterC = c.RegisterA / power(2, combo(i.Operand, c))
		c.InstructionPointer++
	default:
	}
	return c
}

func combo(operand int, c Computer) int {
	switch operand {
	case 4:
		return c.RegisterA
	case 5:
		return c.RegisterB
	case 6:
		return c.RegisterC
	case 7:
		panic("Invalid operand")
	default:
		return operand
	}
}

func literal(operand int) int {
	return operand
}

func power(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func findSolution(instr []Instruction, expected []string, a, i int) int {
	var comp = Computer{RegisterA: a}
	for comp.InstructionPointer < len(instr) {
		comp = executeInstruction(comp, instr[comp.InstructionPointer])
	}

	if comp.Output == strings.Join(expected, ",") {
		return a
	}

	if i == 0 || comp.Output == strings.Join(expected[len(expected)-i:], ",") {
		for j := range 8 {
			na := findSolution(instr, expected, (a<<3)+j, i+1)
			if na > 0 {
				return na
			}
		}
	}

	return 0
}
