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
	l := length(g, start, next)
	fmt.Println("Part 1:", (l+1)/2)
}

type pos struct{ x, y int }

func (p pos) north() pos { return pos{p.x, p.y - 1} }
func (p pos) west() pos  { return pos{p.x - 1, p.y} }
func (p pos) south() pos { return pos{p.x, p.y + 1} }
func (p pos) east() pos  { return pos{p.x + 1, p.y} }

type grid [][]byte

func (g grid) at(p pos) byte { return g[p.y][p.x] }

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
	north := p.north()
	if north.y >= 0 && (g.at(north) == '|' || g.at(north) == '7' || g.at(north) == 'F') {
		return north
	}
	west := p.west()
	if west.x >= 0 && (g.at(west) == '-' || g.at(west) == 'L' || g.at(west) == 'F') {
		return west
	}
	south := p.south()
	if south.y < len(g) && (g.at(south) == '|' || g.at(south) == 'L' || g.at(south) == 'J') {
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

func length(g grid, start, next pos) int {
	prev := start
	count := 0
	for next != start {
		count++
		prev, next = next, g.next(prev, next)
	}
	return count
}
