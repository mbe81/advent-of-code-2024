package day25

import (
	"fmt"
	"time"
)

func part1(input []string) {
	t := time.Now()

	keys, locks := parseInput(input)
	var fits int
	for _, k := range keys {
		for _, l := range locks {
			if k.Pin1+l.Pin1 <= 5 && k.Pin2+l.Pin2 <= 5 && k.Pin3+l.Pin3 <= 5 && k.Pin4+l.Pin4 <= 5 && k.Pin5+l.Pin5 <= 5 {
				fits++
			}
		}
	}

	fmt.Println("Result day 25, part 1:", fits, "- duration:", time.Since(t))
}

func part2(input []string) {
	t := time.Now()

	fmt.Println("Result day 25, part 2:", 0, "- duration:", time.Since(t))
}

type Key struct {
	Pin1 int
	Pin2 int
	Pin3 int
	Pin4 int
	Pin5 int
}

type Lock struct {
	Pin1 int
	Pin2 int
	Pin3 int
	Pin4 int
	Pin5 int
}

func parseInput(input []string) ([]Key, []Lock) {
	var keys []Key
	var locks []Lock

	for i := 0; i < len(input); i++ {
		if input[i] == "#####" {
			var key Key
			for j := 1; j <= 5; j++ {
				if input[i+j][0] == '#' {
					key.Pin1 += 1
				}
				if input[i+j][1] == '#' {
					key.Pin2 += 1
				}
				if input[i+j][2] == '#' {
					key.Pin3 += 1
				}
				if input[i+j][3] == '#' {
					key.Pin4 += 1
				}
				if input[i+j][4] == '#' {
					key.Pin5 += 1
				}
			}
			keys = append(keys, key)
			i += 7
		} else if input[i] == "....." {
			var lock Lock
			for j := 5; j >= 1; j-- {
				if input[i+j][0] == '#' {
					lock.Pin1 += 1
				}
				if input[i+j][1] == '#' {
					lock.Pin2 += 1
				}
				if input[i+j][2] == '#' {
					lock.Pin3 += 1
				}
				if input[i+j][3] == '#' {
					lock.Pin4 += 1
				}
				if input[i+j][4] == '#' {
					lock.Pin5 += 1
				}
			}
			locks = append(locks, lock)
			i += 7
		}
	}
	return keys, locks
}
