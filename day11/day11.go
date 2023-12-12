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
	g := grid{g: bytes.Split(byts, []byte{'\n'})}
	g.expand()
	fmt.Println("Part 1:", sumDistances(g.galaxies(2)))
	fmt.Println("Part 2:", sumDistances(g.galaxies(1000000)))
}

type pos struct{ x, y int }

func abs(a int) int { return max(a, -a) }

func distance(a, b pos) int {
	xDist := abs(b.x - a.x)
	yDist := abs(b.y - a.y)
	return xDist + yDist
}

type grid struct {
	g     [][]byte
	xExps []int
	yExps []int
}

func (g *grid) expand() {
horizontalloop:
	for x := 0; x < len(g.g[0]); x++ {
		for y := range g.g {
			if g.g[y][x] != '.' {
				continue horizontalloop
			}
		}
		g.xExps = append(g.xExps, x)
	}

verticalloop:
	for y := 0; y < len(g.g); y++ {
		for x := range g.g[y] {
			if g.g[y][x] != '.' {
				continue verticalloop
			}
		}
		g.yExps = append(g.yExps, y)
	}
}

func (g grid) galaxies(expFactor int) []pos {
	var gs []pos
	yExpCount := 0
	for y, row := range g.g {
		if yExpCount < len(g.yExps) && g.yExps[yExpCount] == y {
			yExpCount++
			continue
		}
		xExpCount := 0
		for x, cell := range row {
			if xExpCount < len(g.xExps) && g.xExps[xExpCount] == x {
				xExpCount++
				continue
			}
			if cell == '#' {
				gs = append(gs, pos{
					// expFactor-1 because otherwise we over count empty rows and columns
					x + xExpCount*(expFactor-1),
					y + yExpCount*(expFactor-1)})
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
