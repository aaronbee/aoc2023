package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	byts, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(byts), "\n")
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}

type loc struct {
	x, y int
}

const symbols = "#$%&*+-/=@"

func part1(data []string) int {
	parts := map[loc]int{}
	for y, l := range data {
		i := 0
		for {
			x := strings.IndexAny(l[i:], symbols)
			if x == -1 {
				break
			}
			x += i
			addPartsNear(data, parts, x, y)
			i = x + 1
		}
	}
	sum := 0
	for _, part := range parts {
		sum += part
	}
	return sum
}

func addPartsNear(data []string, parts map[loc]int, x, y int) {
	addPart(data, parts, x-1, y-1)
	addPart(data, parts, x-1, y)
	addPart(data, parts, x-1, y+1)
	addPart(data, parts, x, y-1)
	addPart(data, parts, x, y+1)
	addPart(data, parts, x+1, y-1)
	addPart(data, parts, x+1, y)
	addPart(data, parts, x+1, y+1)
}

func isNum(b byte) bool { return '0' <= b && b <= '9' }

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func partBounds(s string, i int) (beg, end int) {
	beg = strings.LastIndexAny(s[:i], symbols+".") + 1
	end = strings.IndexAny(s[i+1:], symbols+".")
	if end == -1 {
		end = len(s)
	} else {
		end += i + 1
	}
	return
}

func addPart(data []string, parts map[loc]int, x, y int) {
	if x < 0 || y < 0 {
		return
	}
	if y >= len(data) || x >= len(data[y]) {
		return
	}
	if !isNum(data[y][x]) {
		return
	}
	beg, end := partBounds(data[y], x)
	parts[loc{beg, y}] = atoi(data[y][beg:end])
}

func part2(data []string) int {
	sum := 0
	for y, l := range data {
		i := 0
		for {
			x := strings.IndexByte(l[i:], '*')
			if x == -1 {
				break
			}
			x += i
			sum += gearRatio(data, x, y)
			i = x + 1
		}
	}
	return sum
}

func gearRatio(data []string, x, y int) int {
	parts := make(map[loc]int)
	addPartsNear(data, parts, x, y)
	if len(parts) != 2 {
		return 0
	}
	ratio := 1
	for _, part := range parts {
		ratio *= part
	}
	return ratio
}
