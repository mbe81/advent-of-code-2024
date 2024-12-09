package day09

import (
	"fmt"
	"time"

	"github.com/mbe81/advent-of-code-2024/lib/convert"
)

func part1(input []string) {
	t := time.Now()

	compactedBlocks := compactPart1(parseInput(input[0]))
	checkSum := calculateChecksum(compactedBlocks)

	fmt.Println("Result day 9, part 1:", checkSum, "- duration:", time.Since(t))
}

func part2(input []string) {
	t := time.Now()

	compactedBlocks := compactPart2(parseInput(input[0]))
	checkSum := calculateChecksum(compactedBlocks)

	fmt.Println("Result day 9, part 2:", checkSum, "- duration:", time.Since(t))
}

func parseInput(diskMap string) []int {
	blocks := make([]int, 0)
	var id int
	for i := 0; i < len(diskMap); i++ {
		entry := convert.StringToInt(diskMap[i : i+1])
		if i%2 == 0 {
			for j := 0; j < entry; j++ {
				blocks = append(blocks, id)
			}
			id++
		} else {
			for j := 0; j < entry; j++ {
				blocks = append(blocks, -1)
			}
		}
	}
	return blocks
}

func compactPart1(blocks []int) []int {
	var blockMoved bool
	for i := len(blocks) - 1; i > 0; i-- {
		blockMoved = false
		for j := 0; j < i; j++ {
			if blocks[j] == -1 {
				blocks[j], blocks[i] = blocks[i], blocks[j]
				blockMoved = true
				break
			}
		}
		if !blockMoved {
			break
		}
	}
	return blocks
}

func compactPart2(blocks []int) []int {
	var freePosition int
	for i := len(blocks) - 1; i > 0; i-- {
		if blocks[i] == -1 {
			continue
		}
		length := 1
		for j := i - 1; j >= 0; j-- {
			if blocks[j] != blocks[i] {
				break
			}
			length++
		}
		freePosition = -1
		for j := 0; j < i; j++ {
			if blocks[j] == -1 {
				freePosition = j
				for k := 1; k < length; k++ {
					if blocks[j+k] != -1 {
						freePosition = -1
						continue
					}
				}
				if freePosition > 0 {
					break
				}
			}
		}
		if freePosition > 0 {
			for k := 0; k < length; k++ {
				blocks[freePosition+k], blocks[i-k] = blocks[i-k], blocks[freePosition+k]
			}
		}
		i = i - length + 1
	}
	return blocks
}

func calculateChecksum(blocks []int) int {
	var checkSum int
	for i := range blocks {
		if blocks[i] != -1 {
			checkSum += i * blocks[i]
		}
	}
	return checkSum
}
