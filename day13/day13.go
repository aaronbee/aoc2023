package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	byts, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	var part1, part2 int
	for _, byts := range bytes.Split(byts, []byte("\n\n")) {
		g := grid(bytes.Split(byts, []byte("\n")))
		if v, ok := g.vSplit(0); ok {
			part1 += v
		} else {
			v, ok := g.hSplit(0)
			if !ok {
				panic("no mirrors")
			}
			part1 += 100 * v
		}
		if v, ok := g.vSplit(1); ok {
			part2 += v
		} else {
			v, ok := g.hSplit(1)
			if !ok {
				panic("no mirrors")
			}
			part2 += 100 * v
		}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

type grid [][]byte

func (g grid) vSplit(allow int) (int, bool) {
	for x := 1; x < len(g[0]); x++ {
		if g.vSmudges(x) == allow {
			return x, true
		}
	}
	return 0, false
}

func (g grid) hSplit(allow int) (int, bool) {
	for y := 1; y < len(g); y++ {
		if g.hSmudges(y) == allow {
			return y, true
		}
	}
	return 0, false
}

func (g grid) vSmudges(x int) int {
	var smudges int
	for _, row := range g {
		l := x - 1
		r := x
		for l >= 0 && r < len(row) {
			if row[l] != row[r] {
				smudges++
			}
			l--
			r++
		}
	}
	return smudges
}

func (g grid) hSmudges(y int) int {
	var smudges int
	a := y - 1
	b := y
	for a >= 0 && b < len(g) {
		for i := 0; i < len(g[a]); i++ {
			if g[a][i] != g[b][i] {
				smudges++
			}
		}
		a--
		b++
	}
	return smudges
}
