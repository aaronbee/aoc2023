package main

import (
	"bufio"
	"fmt"
	"os"
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
	i := 0
	for s.Scan() {
		i++
		springs, cs, ok := strings.Cut(s.Text(), " ")
		if !ok {
			panic("bad line")
		}
		var counts []int
		for _, c := range strings.Split(cs, ",") {
			counts = append(counts, aoc2023.Atoi(c))
		}
		r := row{springs: springs, counts: counts}
		part1 += r.arrangements()
		part2 += r.arrangements2()
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

type cacheKey struct {
	s  string
	cs int
}

type row struct {
	springs string
	counts  []int
}

func (r *row) arrangements() int {
	return matchCount(r.springs, r.counts, make(map[cacheKey]int))
}

func (r *row) arrangements2() int {
	var buf strings.Builder
	counts := make([]int, len(r.counts)*5)
	for i := 0; i < 5; i++ {
		if i != 0 {
			buf.WriteByte('?')
		}
		buf.WriteString(r.springs)

		copy(counts[i*len(r.counts):], r.counts)
	}
	return matchCount(buf.String(), counts, make(map[cacheKey]int))
}

func length(cs []int) int {
	l := len(cs) - 1
	for _, c := range cs {
		l += c
	}
	return l
}

func valid(s string, run, offset int) bool {
	if offset != 0 && s[offset-1] == '#' {
		// Need a blank space before run
		return false
	}
	if run+offset > len(s) {
		return false
	}
	if run+offset < len(s) && s[run+offset] == '#' {
		// Need a blank space after run
		return false
	}
	for _, c := range s[offset : offset+run] {
		if c == '.' {
			return false
		}
	}
	return true
}

func matchCount(s string, cs []int, cache map[cacheKey]int) int {
	if len(cs) == 0 {
		if strings.Contains(s, "#") {
			// invalid, didn't cover all the '#
			return 0
		}
		return 1
	}
	l := length(cs)
	if l > len(s) {
		return 0
	}
	if v, ok := cache[cacheKey{s, len(cs)}]; ok {
		return v
	}
	var count int
	run, rest := cs[0], cs[1:]
	for i := 0; i <= len(s)-l; i++ {
		if valid(s, run, i) {
			if run+i == len(s) {
				count++
				break
			}
			count += matchCount(s[run+i+1:], rest, cache)
		}
		if s[i] == '#' {
			// can't skip over any '#'
			break
		}
	}
	cache[cacheKey{s, len(cs)}] = count
	return count
}
