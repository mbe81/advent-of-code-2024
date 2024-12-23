package day23

import (
	"fmt"
	"slices"
	"sort"
	"strings"
	"time"
)

func part1(input []string) {
	t := time.Now()

	conns := make(map[string]map[string]bool)
	for _, line := range input {
		comp := strings.Split(line, "-")
		if conns[comp[0]] == nil {
			conns[comp[0]] = make(map[string]bool)
		}
		if conns[comp[1]] == nil {
			conns[comp[1]] = make(map[string]bool)
		}
		conns[comp[0]][comp[1]] = true
		conns[comp[1]][comp[0]] = true
	}

	type set [3]string
	sets := make(map[set]bool)
	for comp1 := range conns {
		for comp2 := range conns[comp1] {
			for comp3 := range conns[comp1] {
				if comp1 == comp2 || comp1 == comp3 || comp2 >= comp3 {
					continue
				}
				if conns[comp2][comp3] {
					comps := []string{comp1, comp2, comp3}
					sort.Strings(comps)
					sets[set{comps[0], comps[1], comps[2]}] = true
				}
			}
		}
	}

	count := 0
	for comps := range sets {
		for _, comp := range comps {
			if comp[0] == 't' {
				count++
				break
			}
		}
	}

	fmt.Println("Result day 22, part 1:", count, "- duration:", time.Since(t))
}

func part2(input []string) {
	t := time.Now()

	var comps = make([]string, 0)
	for _, line := range input {
		split := strings.Split(line, "-")
		if !slices.Contains(comps, split[0]) {
			comps = append(comps, split[0])
		}
		if !slices.Contains(comps, split[1]) {
			comps = append(comps, split[1])
		}
	}
	slices.Sort(comps)

	conns := make(map[string]map[string]bool)
	for _, line := range input {
		comp := strings.Split(line, "-")
		if conns[comp[0]] == nil {
			conns[comp[0]] = make(map[string]bool)
		}
		if conns[comp[1]] == nil {
			conns[comp[1]] = make(map[string]bool)
		}
		conns[comp[0]][comp[1]] = true
		conns[comp[1]][comp[0]] = true
	}

	findCliques([]string{}, comps, conns)

	var longestClique []string
	var maxLength int
	for _, clique := range cliques {
		if len(clique) > maxLength {
			longestClique = clique
			maxLength = len(clique)
		}
	}
	sort.Strings(longestClique)

	fmt.Println("Result day 22, part 2:", "'"+strings.Join(longestClique, ",")+"'", "- duration:", time.Since(t))
}

var cliques = make([][]string, 0)

func findCliques(clique []string, comps []string, conns map[string]map[string]bool) {
	for _, comp := range comps {
		if len(clique) > 0 && comp < clique[len(clique)-1] {
			continue
		}
		if slices.Contains(clique, comp) {
			continue
		}
		extend := true
		for _, c := range clique {
			if !conns[comp][c] {
				extend = false
				break
			}
		}
		if extend {
			newClique := slices.Clone(clique)
			newClique = append(newClique, comp)
			cliques = append(cliques, newClique)
			if len(comps[1:]) > 0 {
				findCliques(newClique, comps[1:], conns)
			}
		}
	}
}
