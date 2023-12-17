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
	fmt.Println("Part 1:", energize(g))
}

func energize(g aoc2023.Grid2D) int {
	// lit is map of locations already visited and directions from the location.
	lit := make(map[aoc2023.Pos2D][4]bool)
	walk(g, lit, aoc2023.Pos2D{X: 0, Y: 0}, aoc2023.East)
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
