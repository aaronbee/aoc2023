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
}

type grid [][]byte

func (g grid) col(x int) column { return column{g: g, x: x} }
func (g grid) row(y int) row    { return row(g[y]) }

func (g grid) print() {
	for _, row := range g {
		fmt.Printf("%s\n", row)
	}
}

func (g grid) tiltN() {
	for x := range g[0] {
		g.col(x).tiltN()
	}
}

func (g grid) load() int {
	var l int
	for y, row := range g {
		l += bytes.Count(row, []byte("O")) * (len(g) - y)
	}
	return l
}

type column struct {
	g grid
	x int
}

func (c column) iter(f func(y int, v byte) bool) {
	for y, row := range c.g {
		if !f(y, row[c.x]) {
			return
		}
	}
}
func (c column) set(y int, v byte) { c.g[y][c.x] = v }

func (c column) tiltN() {
	stopLoc := 0
	c.iter(func(y int, v byte) bool {
		switch v {
		case '#':
			stopLoc = y + 1
		case 'O':
			if stopLoc != y {
				c.set(stopLoc, 'O')
				c.set(y, '.')
			}
			stopLoc++
		}
		return true
	})
}

type row []byte

func (r row) iter(f func(x int, v byte) bool) {
	for x, v := range r {
		if !f(x, v) {
			return
		}
	}
}
func (r row) set(x int, v byte) { r[x] = v }
