package main

import (
	"bytes"
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
	p := aoc2023.Pos2D{X: 0, Y: 0}
	best := energize(g, p, aoc2023.East)
	fmt.Println("Part 1:", best)
	for {
		p = p.South()
		if p.Y == len(g) {
			p = p.North()
			break
		}
		if c := energize(g, p, aoc2023.East); c > best {
			best = c
		}
	}
	for {
		p = p.East()
		if p.X == len(g[0]) {
			p = p.West()
			break
		}
		if c := energize(g, p, aoc2023.North); c > best {
			best = c
		}
	}
	for {
		p = p.North()
		if p.Y < 0 {
			p = p.South()
			break
		}
		if c := energize(g, p, aoc2023.West); c > best {
			best = c
		}
	}
	for {
		p = p.West()
		if p.X < 0 {
			break
		}
		if c := energize(g, p, aoc2023.South); c > best {
			best = c
		}
	}
	fmt.Println("Part 2:", best)
}

func energize(g aoc2023.Grid2D, p aoc2023.Pos2D, d aoc2023.Dir) int {
	// lit is map of locations already visited and directions from the location.
	lit := make(map[aoc2023.Pos2D][4]bool)
	walk(g, lit, p, d)
	return len(lit)
}

func cache(lit map[aoc2023.Pos2D][4]bool, p aoc2023.Pos2D, d aoc2023.Dir) bool {
	ds := lit[p]
	if ds[d] {
		return true
	}
	ds[d] = true
	lit[p] = ds
	return false
}

func walk(g aoc2023.Grid2D, lit map[aoc2023.Pos2D][4]bool, p aoc2023.Pos2D, d aoc2023.Dir) {
	g.Iter(p, d, func(p aoc2023.Pos2D, v byte) bool {
		if cache(lit, p, d) {
			return false
		}
		switch v {
		case '.':
		case '|':
			switch d {
			case aoc2023.East, aoc2023.West:
				walk(g, lit, p.North(), aoc2023.North)
				walk(g, lit, p.South(), aoc2023.South)
				return false
			}
		case '-':
			switch d {
			case aoc2023.North, aoc2023.South:
				walk(g, lit, p.West(), aoc2023.West)
				walk(g, lit, p.East(), aoc2023.East)
				return false
			}
		case '/':
			switch d {
			case aoc2023.North:
				walk(g, lit, p.East(), aoc2023.East)
			case aoc2023.East:
				walk(g, lit, p.North(), aoc2023.North)
			case aoc2023.South:
				walk(g, lit, p.West(), aoc2023.West)
			case aoc2023.West:
				walk(g, lit, p.South(), aoc2023.South)
			}
			return false
		case '\\':
			switch d {
			case aoc2023.North:
				walk(g, lit, p.West(), aoc2023.West)
			case aoc2023.East:
				walk(g, lit, p.South(), aoc2023.South)
			case aoc2023.South:
				walk(g, lit, p.East(), aoc2023.East)
			case aoc2023.West:
				walk(g, lit, p.North(), aoc2023.North)
			}
			return false
		default:
			panic(fmt.Errorf("unexpected v at %v: %q", p, rune(v)))
		}
		return true
	})
}
