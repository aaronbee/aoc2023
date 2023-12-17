package main

import (
	"bytes"
	"cmp"
	"fmt"
	"os"

	"github.com/aaronbee/aoc2023"
)

func main() {
	byts, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	g := aoc2023.Grid2D(bytes.Split(byts, []byte("\n")))
	fmt.Println("Part 1:", walk(g))
}

type pos struct {
	p     aoc2023.Pos2D
	d     aoc2023.Dir
	count int
}

type step struct {
	cost int
	p    pos
}

func (s *step) next(g aoc2023.Grid2D, yield func(*step)) {
	n := func(d aoc2023.Dir) {
		n := step{
			p: pos{
				p:     s.p.p.Step(d),
				d:     d,
				count: 1,
			},
		}
		cost, ok := g.Get(n.p.p)
		if !ok {
			return
		}
		n.cost = s.cost + int(cost-'0')
		if s.p.d == d {
			n.p.count = s.p.count + 1
		}
		yield(&n)
	}
	if s.p.count < 3 {
		n(s.p.d)
	}
	switch s.p.d {
	case aoc2023.North:
		n(aoc2023.East)
		n(aoc2023.West)
	case aoc2023.East:
		n(aoc2023.North)
		n(aoc2023.South)
	case aoc2023.South:
		n(aoc2023.East)
		n(aoc2023.West)
	case aoc2023.West:
		n(aoc2023.North)
		n(aoc2023.South)
	}
}

func walk(g aoc2023.Grid2D) int {
	seen := map[pos]int{}
	q := aoc2023.NewHeap(func(a, b *step) int {
		aV := a.p.p.X + a.p.p.Y - a.cost
		bV := b.p.p.X + b.p.p.Y - b.cost
		return cmp.Compare(aV, bV)
	})
	q.Push(&step{p: pos{d: aoc2023.East}})
	i := 0
	for q.Len() > 0 {
		i++
		s := q.Pop()
		if s.p.p.Y == len(g)-1 && s.p.p.X == len(g[s.p.p.Y])-1 {
			fmt.Printf("found end in %d iterations\n", i)
			return s.cost
		}
		if cost, ok := seen[s.p]; ok && cost <= s.cost {
			continue
		}
		seen[s.p] = s.cost
		s.next(g, func(s *step) { q.Push(s) })
	}
	panic("path to end not found")
}
