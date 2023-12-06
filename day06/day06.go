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
	s.Scan()
	var times []int
	timesLine := strings.TrimPrefix(s.Text(), "Time:")
	aoc2023.IntFieldsIter(timesLine, func(i int) { times = append(times, i) })
	s.Scan()
	var distances []int
	distancesLine := strings.TrimPrefix(s.Text(), "Distance:")
	aoc2023.IntFieldsIter(distancesLine, func(i int) { distances = append(distances, i) })

	fmt.Println("Part 1:", part1(times, distances))

	timesLine = strings.Replace(timesLine, " ", "", -1)
	distancesLine = strings.Replace(distancesLine, " ", "", -1)
	t, err := strconv.Atoi(timesLine)
	if err != nil {
		panic(err)
	}
	d, err := strconv.Atoi(distancesLine)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 2:", doBetterThan(t, d))
}

func part1(times, distances []int) int {
	prod := 1
	for i := 0; i < len(times); i++ {
		prod *= doBetterThan(times[i], distances[i])
	}
	return prod
}

func doBetterThan(t, d int) int {
	sum := 0
	for i := 1; i < t; i++ {
		if distance(t, i) > d {
			sum++
		}
	}
	return sum
}

func distance(totalTime int, timeHeld int) int {
	if timeHeld >= totalTime {
		return 0
	}
	return timeHeld * (totalTime - timeHeld)
}
