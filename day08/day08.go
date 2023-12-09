package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	fmt.Println("Part 1:", part1(dirs, nodes))
}

func part1(dirs string, nodes map[string]node) int {
	count := 0
	node, ok := nodes["AAA"]
	if !ok {
		panic(fmt.Errorf("cant find node AAA"))
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
		if next == "ZZZ" {
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
