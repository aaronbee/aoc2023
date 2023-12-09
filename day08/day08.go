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
	s.Scan()
	dirs := s.Text()
	s.Scan()
	nodes := make(map[string]node)
	for s.Scan() {
		n, lr, ok := strings.Cut(s.Text(), " = ")
		if !ok {
			panic(fmt.Errorf("unexpected line: %q", s.Text()))
		}
		l, r, ok := strings.Cut(strings.Trim(lr, "()"), ", ")
		if !ok {
			panic(fmt.Errorf("unexpected rhs: %q", lr))
		}
		nodes[n] = node{l, r}
	}
	c := walk(dirs, nodes, "AAA", func(s string) bool { return s == "ZZZ" })
	fmt.Println("Part 1:", c)

	fmt.Println("Part 2:", part2(dirs, nodes))
}

func walk(dirs string, nodes map[string]node, start string, end func(string) bool) int {
	count := 0
	node, ok := nodes[start]
	if !ok {
		panic(fmt.Errorf("cant find node %q", start))
	}
	for {
		dir := dirs[count%len(dirs)]
		var next string
		switch dir {
		case 'L':
			next = node.l
		case 'R':
			next = node.r
		default:
			panic(fmt.Errorf("unexpected dir: %q", dir))
		}
		count++
		if end(next) {
			return count
		}
		node, ok = nodes[next]
		if !ok {
			panic(fmt.Errorf("cant find node %q", next))
		}
	}
}

type node struct {
	l, r string
}

func part2(dirs string, nodes map[string]node) int {
	var stepsToZ []int
	for n := range nodes {
		if strings.HasSuffix(n, "A") {
			pos := walk(dirs, nodes, n, func(s string) bool {
				return strings.HasSuffix(s, "Z")
			})
			stepsToZ = append(stepsToZ, pos)
		}
	}
	return aoc2023.GCM(stepsToZ)
}
