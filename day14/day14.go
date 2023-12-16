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
	g.tiltN()
	fmt.Println("Part 1:", g.load())
	g.tiltW()
	g.tiltS()
	g.tiltE()
	cur, l := cycleLength(g, make(map[string]int))
	for cur%l != 1000000000%l {
		g.cycle()
		cur++
	}
	fmt.Println("Part 2:", g.load())
}

func cycleLength(g grid, cache map[string]int) (int, int) {
	c := 1
	k := bytes.Join(g, nil)
	cache[string(k)] = c
	for {
		g.cycle()
		c++
		k := bytes.Join(g, nil)
		if cycle, ok := cache[string(k)]; ok {
			return c, c - cycle
		}
		cache[string(k)] = c
	}
}

type run interface {
	iter(func(i int, v byte))
	set(i int, v byte)
}

func tilt(r run) {
	stopLoc := 0
	r.iter(func(i int, v byte) {
		switch v {
		case '#':
			stopLoc = i + 1
		case 'O':
			if stopLoc != i {
				r.set(stopLoc, 'O')
				r.set(i, '.')
			}
			stopLoc++
		}
	})
}

type grid [][]byte

func (g grid) south(x int) southRun { return southRun{g: g, x: x} }
func (g grid) east(y int) eastRun   { return eastRun(g[y]) }
func (g grid) north(x int) northRun { return northRun{g: g, x: x} }
func (g grid) west(y int) westRun   { return westRun(g[y]) }

func (g grid) cycle() {
	g.tiltN()
	g.tiltW()
	g.tiltS()
	g.tiltE()
}

func (g grid) tiltN() {
	for x := range g[0] {
		tilt(g.south(x))
	}
}
func (g grid) tiltW() {
	for y := range g {
		tilt(g.east(y))
	}
}
func (g grid) tiltS() {
	for x := range g[0] {
		tilt(g.north(x))
	}
}
func (g grid) tiltE() {
	for y := range g {
		tilt(g.west(y))
	}
}

func (g grid) load() int {
	var l int
	for y, row := range g {
		l += bytes.Count(row, []byte("O")) * (len(g) - y)
	}
	return l
}

type southRun struct {
	g grid
	x int
}

func (r southRun) iter(f func(y int, v byte)) {
	for y, row := range r.g {
		f(y, row[r.x])
	}
}
func (r southRun) set(y int, v byte) { r.g[y][r.x] = v }

type northRun struct {
	g grid
	x int
}

func (r northRun) iter(f func(i int, v byte)) {
	for y := len(r.g) - 1; y >= 0; y-- {
		i := len(r.g) - 1 - y
		f(i, r.g[y][r.x])
	}
}
func (r northRun) set(i int, v byte) {
	y := len(r.g) - 1 - i
	r.g[y][r.x] = v
}

type eastRun []byte

func (r eastRun) iter(f func(x int, v byte)) {
	for x, v := range r {
		f(x, v)
	}
}
func (r eastRun) set(x int, v byte) { r[x] = v }

type westRun []byte

func (r westRun) iter(f func(i int, v byte)) {
	for x := len(r) - 1; x >= 0; x-- {
		i := len(r) - 1 - x
		f(i, r[x])
	}
}
func (r westRun) set(i int, v byte) {
	x := len(r) - 1 - i
	r[x] = v
}
