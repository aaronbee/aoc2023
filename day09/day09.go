package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"

	"github.com/aaronbee/aoc2023"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	var lines [][]int
	for s.Scan() {
		fs := bytes.Fields(s.Bytes())
		nums := slices.Grow([]int(nil), len(fs))
		for _, f := range fs {
			nums = append(nums, aoc2023.Atoi(string(f)))
		}
		lines = append(lines, nums)
	}
	fmt.Println("Part 1:", part1(lines))
}

func part1(lines [][]int) int {
	sum := 0
	for _, l := range lines {
		sum += next(l)
	}
	return sum
}

func allEqual(s []int) bool {
	if len(s) == 0 {
		return true
	}
	v := s[0]
	for _, vv := range s[1:] {
		if v != vv {
			return false
		}
	}
	return true
}

func last[S ~[]E, E any](s S) E {
	return s[len(s)-1]
}

func next(l []int) int {
	nums := append([][]int(nil), l)
	for {
		cur := last(nums)
		next := make([]int, len(cur)-1, len(cur))
		for i := range next {
			next[i] = cur[i+1] - cur[i]
		}
		if allEqual(next) {
			next = append(next, next[0])
			nums = append(nums, next)
			break
		}
		nums = append(nums, next)
	}
	for j := len(nums) - 1; j > 0; j-- {
		diff := last(nums[j])
		nums[j-1] = append(nums[j-1], last(nums[j-1])+diff)
	}

	return last(nums[0])
}
