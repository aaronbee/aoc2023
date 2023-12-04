package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aaronbee/aoc2023"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)

	var part1, part2 int

	for s.Scan() {
		part1 += evalGame1(s.Text())
		part2 += evalGame2(s.Text())
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func evalGame1(game string) int {
	limits := counter{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	g, sets, ok := strings.Cut(game, ":")
	if !ok {
		panic(fmt.Errorf("unexpected game: %q", game))
	}
	id := atoi(strings.TrimPrefix(g, "Game "))

	for _, set := range strings.Split(sets, ";") {
		cntr := setToCounter(set)
		if !cntr.le(limits) {
			return 0
		}
	}
	return id
}

func evalGame2(game string) int {
	_, sets, ok := strings.Cut(game, ":")
	if !ok {
		panic(fmt.Errorf("unexpected game: %q", game))
	}
	acc := make(counter)
	for _, set := range strings.Split(sets, ";") {
		cntr := setToCounter(set)
		acc.union(cntr)
	}
	return acc.power()
}

type counter map[string]int

func setToCounter(set string) counter {
	cntr := make(counter)
	for _, cubes := range strings.Split(set, ",") {
		count, color, ok := strings.Cut(strings.TrimSpace(cubes), " ")
		if !ok {
			panic(fmt.Errorf("unexpected play: %q", set))
		}
		cntr[color] += atoi(count)
	}
	return cntr
}

func (cntr counter) le(o counter) bool {
	for k, c := range cntr {
		if c > o[k] {
			return false
		}
	}
	return true
}

func (cntr counter) union(o counter) {
	for k, c := range o {
		cntr[k] = max(cntr[k], c)
	}
}

func (cntr counter) power() int {
	return aoc2023.ProdMapVal(cntr)
}
