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
	var part1, part2 int
	for s.Scan() {
		fs := bytes.Fields(s.Bytes())
		nums := slices.Grow([]int(nil), len(fs))
		for _, f := range fs {
			nums = append(nums, aoc2023.Atoi(string(f)))
		}
		ns := expand(nums)
		part1 += next(ns)
		part2 += prev(ns)
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
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

func expand(l []int) [][]int {
	nums := append([][]int(nil), l)
	for {
		cur := last(nums)
		next := make([]int, len(cur)-1)
		for i := range next {
			next[i] = cur[i+1] - cur[i]
		}
		nums = append(nums, next)
		if allEqual(next) {
			break
		}
	}
	return nums
}

func next(nums [][]int) int {
	diff := last(nums)[0]
	for j := len(nums) - 2; j >= 0; j-- {
		diff += last(nums[j])
	}
	return diff
}

func prev(nums [][]int) int {
	diff := last(nums)[0]
	for j := len(nums) - 2; j >= 0; j-- {
		diff = nums[j][0] - diff
	}
	return diff
}
