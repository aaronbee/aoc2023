package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	byts, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	g := grid(bytes.Split(byts, []byte{'\n'}))
	g = g.expand()
	gs := g.galaxies()
	fmt.Println("Part 1:", sumDistances(gs))
}

type pos struct{ x, y int }

func abs(a int) int { return max(a, -a) }

func distance(a, b pos) int {
	xDist := abs(b.x - a.x)
	yDist := abs(b.y - a.y)
	return xDist + yDist
}

type grid [][]byte

func (g grid) expand() grid {
horizontalloop:
	for x := 0; x < len(g[0]); x++ {
		for y := range g {
			if g[y][x] != '.' {
				continue horizontalloop
			}
		}
		for y := range g {
			g[y] = slices.Insert(g[y], x, '.')
		}
		x++
	}

verticalloop:
	for y := 0; y < len(g); y++ {
		for x := range g[y] {
			if g[y][x] != '.' {
				continue verticalloop
			}
		}
		newRow := make([]byte, len(g[y]))
		copy(newRow, g[y])
		g = slices.Insert(g, y, newRow)
		y++
	}

	return g
}

func (g grid) galaxies() []pos {
	var gs []pos
	for y, row := range g {
		for x, cell := range row {
			if cell == '#' {
				gs = append(gs, pos{x, y})
			}
		}
	}
	return gs
}

func sumDistances(galaxies []pos) int {
	var sum int
	for i, a := range galaxies {
		for _, b := range galaxies[i+1:] {
			sum += distance(a, b)
		}
	}
	return sum
}
