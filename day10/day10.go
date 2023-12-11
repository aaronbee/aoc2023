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
	g := grid(bytes.Split(byts, []byte("\n")))
	start := g.start()
	next := g.startNext(start)
	ng := make(grid, len(g))
	for i := range ng {
		ng[i] = make([]byte, len(g[i]))
		for j := range ng[i] {
			ng[i][j] = '.'
		}
	}
	l := g.length(ng, start, next)
	fmt.Println("Part 1:", (l+1)/2)

	ng.color()
	fmt.Println("Part 2:", ng.countInside())
}

type pos struct{ x, y int }

func (p pos) north() pos { return pos{p.x, p.y - 1} }
func (p pos) west() pos  { return pos{p.x - 1, p.y} }
func (p pos) south() pos { return pos{p.x, p.y + 1} }
func (p pos) east() pos  { return pos{p.x + 1, p.y} }

type grid [][]byte

func (g grid) at(p pos) byte     { return g[p.y][p.x] }
func (g grid) set(p pos, b byte) { g[p.y][p.x] = b }

func (g grid) start() pos {
	for y, row := range g {
		for x, cell := range row {
			if cell == 'S' {
				return pos{x, y}
			}
		}
	}
	panic("no start")
}

func (g grid) startNext(p pos) pos {
	var (
		n, w, s bool
	)
	north := p.north()
	if north.y >= 0 && (g.at(north) == '|' || g.at(north) == '7' || g.at(north) == 'F') {
		n = true
	}
	west := p.west()
	if west.x >= 0 && (g.at(west) == '-' || g.at(west) == 'L' || g.at(west) == 'F') {
		w = true
	}
	south := p.south()
	if south.y < len(g) && (g.at(south) == '|' || g.at(south) == 'L' || g.at(south) == 'J') {
		s = true
	}
	if n && w {
		g.set(p, 'J')
		return north
	}
	if n && s {
		g.set(p, '|')
		return north
	}
	if n {
		g.set(p, 'L')
		return north
	}
	if w && s {
		g.set(p, '7')
		return west
	}
	if w {
		g.set(p, '-')
		return west
	}
	if s {
		g.set(p, 'F')
		return south
	}
	panic("can't find next")
}

func (g grid) next(prev, cur pos) pos {
	switch g.at(cur) {
	case '|':
		if prev == cur.north() {
			return cur.south()
		} else {
			return cur.north()
		}
	case '-':
		if prev == cur.east() {
			return cur.west()
		} else {
			return cur.east()
		}
	case 'L':
		if prev == cur.north() {
			return cur.east()
		} else {
			return cur.north()
		}
	case 'J':
		if prev == cur.north() {
			return cur.west()
		} else {
			return cur.north()
		}
	case '7':
		if prev == cur.west() {
			return cur.south()
		} else {
			return cur.west()
		}
	case 'F':
		if prev == cur.east() {
			return cur.south()
		} else {
			return cur.east()
		}
	}
	panic(fmt.Errorf("invalid pipe: %q", g.at(cur)))
}

func (g grid) length(ng grid, start, next pos) int {
	prev := start
	count := 0
	for next != start {
		ng.set(prev, g.at(prev))
		count++
		prev, next = next, g.next(prev, next)
	}
	ng.set(prev, g.at(prev))
	return count
}

func (g grid) color() {
	for y := 0; y < len(g); y++ {
		var (
			inside bool
			start  byte // set when on top of a pipe
		)
		for x := 0; x < len(g[y]); x++ {
			p := pos{x, y}
			switch c := g.at(p); c {
			case '.':
				if start != 0 {
					panic("unexpected state")
				}
				if inside {
					g.set(p, 'I')
				} else {
					g.set(p, 'O')
				}
			case '|':
				if start != 0 {
					panic("unexpected state")
				}
				inside = !inside
			case '-':
				if start == 0 {
					panic("unexpected state")
				}
			case 'L', 'F':
				if start != 0 {
					panic("unexpected state")
				}
				start = c
			case '7':
				if start == 'L' {
					inside = !inside
					start = 0
				} else if start == 'F' {
					start = 0
				} else {
					panic("unexpected state: start=%s c=%s")
				}
			case 'J':
				if start == 'F' {
					inside = !inside
					start = 0
				} else if start == 'L' {
					start = 0
				} else {
					panic("unexpected state")
				}
			}
		}
	}
}

func (g grid) countInside() int {
	var count int
	for _, row := range g {
		for _, cell := range row {
			if cell == 'I' {
				count++
			}
		}
	}
	return count
}
