package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/aaronbee/aoc2023"
)

func main() {
	byts, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(byts), "\n")
	counts := make([]int, len(lines))
	for i := range counts {
		counts[i] = 1
	}
	var part1 int
	for n, line := range lines {
		part1 += evalGame1(line)

		score := score2(line)
		count := counts[n]
		for i := 0; i < score && n+i+1 < len(counts); i++ {
			counts[n+i+1] += count
		}
	}
	part2 := aoc2023.SumSlice(counts)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func evalGame1(s string) int {
	_, card, ok := strings.Cut(s, ": ")
	if !ok {
		panic(fmt.Errorf("unexpected line: %q", s))
	}
	winners, plays, ok := strings.Cut(card, " | ")
	if !ok {
		panic(fmt.Errorf("unexpected card: %q", card))
	}
	w := make(map[string]struct{})
	for _, n := range strings.Fields(winners) {
		w[n] = struct{}{}
	}
	score := 0
	for _, n := range strings.Fields(plays) {
		if _, ok := w[n]; ok {
			if score == 0 {
				score = 1
			} else {
				score <<= 1
			}
		}
	}
	return score
}

func score2(s string) int {
	_, card, ok := strings.Cut(s, ": ")
	if !ok {
		panic(fmt.Errorf("unexpected line: %q", s))
	}
	winners, plays, ok := strings.Cut(card, " | ")
	if !ok {
		panic(fmt.Errorf("unexpected card: %q", card))
	}
	w := make(map[string]struct{})
	for _, n := range strings.Fields(winners) {
		w[n] = struct{}{}
	}
	score := 0
	for _, n := range strings.Fields(plays) {
		if _, ok := w[n]; ok {
			score++
		}
	}
	return score
}
