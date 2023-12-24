package aoc2023

import "fmt"

type Pos2D struct{ X, Y int }

func (p Pos2D) North() Pos2D { return Pos2D{p.X, p.Y - 1} }
func (p Pos2D) East() Pos2D  { return Pos2D{p.X + 1, p.Y} }
func (p Pos2D) South() Pos2D { return Pos2D{p.X, p.Y + 1} }
func (p Pos2D) West() Pos2D  { return Pos2D{p.X - 1, p.Y} }

func (p Pos2D) Step(d Dir) Pos2D {
	switch d {
	case North:
		return p.North()
	case East:
		return p.East()
	case South:
		return p.South()
	case West:
		return p.West()
	}
	panic(fmt.Errorf("unexpected dir: %d", d))
}

type Dir int

const (
	North Dir = iota
	East
	South
	West
)

type Grid2D [][]byte

func (g Grid2D) Print() {
	for _, row := range g {
		fmt.Printf("%s\n", row)
	}
}

func (g Grid2D) Get(p Pos2D) (byte, bool) {
	if p.Y < 0 || p.Y >= len(g) {
		return 0, false
	}
	if p.X < 0 || p.X >= len(g[p.Y]) {
		return 0, false
	}
	return g[p.Y][p.X], true
}

func (g Grid2D) Set(p Pos2D, b byte) {
	g[p.Y][p.X] = b
}

func (g Grid2D) Iter(p Pos2D, d Dir, yield func(p Pos2D, v byte) bool) {
	if p.Y < 0 || p.Y >= len(g) {
		return
	}
	if p.X < 0 || p.X >= len(g[p.Y]) {
		return
	}
	switch d {
	case North:
		NorthIter{g, p.X, p.Y}.Iter(yield)
	case East:
		EastIter{g, p.X, p.Y}.Iter(yield)
	case South:
		SouthIter{g, p.X, p.Y}.Iter(yield)
	case West:
		WestIter{g, p.X, p.Y}.Iter(yield)
	default:
		panic(fmt.Errorf("unexpected dir: %d", d))
	}
}

type SouthIter struct {
	g      Grid2D
	x      int
	startY int
}

func (r SouthIter) Iter(yield func(p Pos2D, v byte) bool) {
	for y := r.startY; y < len(r.g); y++ {
		if !yield(Pos2D{r.x, y}, r.g[y][r.x]) {
			break
		}
	}
}

type NorthIter struct {
	g      Grid2D
	x      int
	startY int
}

func (r NorthIter) Iter(yield func(p Pos2D, v byte) bool) {
	for y := r.startY; y >= 0; y-- {
		if !yield(Pos2D{r.x, y}, r.g[y][r.x]) {
			break
		}
	}
}

type EastIter struct {
	g      Grid2D
	startX int
	y      int
}

func (r EastIter) Iter(yield func(p Pos2D, v byte) bool) {
	for x := r.startX; x < len(r.g[r.y]); x++ {
		if !yield(Pos2D{x, r.y}, r.g[r.y][x]) {
			break
		}
	}
}

type WestIter struct {
	g      Grid2D
	startX int
	y      int
}

func (r WestIter) Iter(yield func(p Pos2D, v byte) bool) {
	for x := r.startX; x >= 0; x-- {
		if !yield(Pos2D{x, r.y}, r.g[r.y][x]) {
			break
		}
	}
}
