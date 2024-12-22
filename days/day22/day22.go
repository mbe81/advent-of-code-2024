package day22

import (
	"fmt"
	"time"

	"github.com/mbe81/advent-of-code-2024/lib/convert"
)

func part1(input []string) {
	t := time.Now()

	totalSum := 0
	for _, line := range input {
		secret := convert.StringToInt(line)
		for i := 0; i < 2000; i++ {
			secret = nextSecret(secret)
		}
		totalSum += secret
	}

	fmt.Println("Result day 22, part 1:", totalSum, "- duration:", time.Since(t))
}

func part2(input []string) {
	t := time.Now()

	type Sequence struct {
		s1, s2, s3, s4 int8
	}

	secretSequences := make(map[int]map[Sequence]int)
	for _, line := range input {
		var s Sequence

		initialSecret := convert.StringToInt(line)
		prevSecret, secret := 0, initialSecret

		for i := 0; i < 2000; i++ {
			prevSecret, secret = secret, nextSecret(secret)
			if prevSecret > 0 {
				s.s1, s.s2, s.s3, s.s4 = s.s2, s.s3, s.s4, nextSequence(prevSecret, secret)
			}
			if i > 3 {
				if _, ok := secretSequences[initialSecret]; !ok {
					secretSequences[initialSecret] = make(map[Sequence]int)
				}
				if _, ok := secretSequences[initialSecret][s]; !ok {
					secretSequences[initialSecret][s] = secret % 10
				}
			}
		}
	}

	sequencePrizes := make(map[Sequence]int)
	for _, sequences := range secretSequences {
		for sequence, prize := range sequences {
			if _, ok := sequencePrizes[sequence]; !ok {
				sequencePrizes[sequence] = prize
			} else {
				sequencePrizes[sequence] += prize
			}
		}
	}

	maxPrize := 0
	for _, prize := range sequencePrizes {
		if prize > maxPrize {
			maxPrize = prize
		}
	}

	fmt.Println("Result day 22, part 2:", maxPrize, "- duration:", time.Since(t))
}

func nextSecret(secretNumber int) int {
	secretNumber = pruning(mixing(secretNumber, secretNumber*64))
	secretNumber = pruning(mixing(secretNumber, secretNumber/32))
	secretNumber = pruning(mixing(secretNumber, secretNumber*2048))
	return secretNumber
}

func mixing(secretNumber, value int) int {
	return secretNumber ^ value
}
func pruning(secretNumber int) int {
	return secretNumber % 16777216
}

func nextSequence(prevSecret, secret int) int8 {
	return int8(secret%10 - prevSecret%10)
}
