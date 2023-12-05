package main

import (
	"fmt"
	"os"
	"strings"
	"time"

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
	t := time.Now()
	var part1 int
	for n, line := range lines {
		m := matches(line)
		if m == 0 {
			continue
		}
		part1 += 1 << (m - 1)

		count := counts[n]
		for i := 0; i < m && n+i+1 < len(counts); i++ {
			counts[n+i+1] += count
		}
	}
	part2 := aoc2023.SumSlice(counts)
	fmt.Println("time:", time.Since(t))
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func matches(s string) int {
	_, card, ok := strings.Cut(s, ": ")
	if !ok {
		panic(fmt.Errorf("unexpected line: %q", s))
	}
	winners, plays, ok := strings.Cut(card, " | ")
	if !ok {
		panic(fmt.Errorf("unexpected card: %q", card))
	}
	boolSet := [101]bool{}
	aoc2023.IntFieldsIter(winners, func(v int) {
		boolSet[v] = true
	})
	var count int
	aoc2023.IntFieldsIter(plays, func(v int) {
		if boolSet[v] {
			count++
		}
	})

	return count
}
