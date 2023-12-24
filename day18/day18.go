package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/aaronbee/aoc2023"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	var insts []inst
	var curY, minY, maxY int
	var curX, minX, maxX int
	s := bufio.NewScanner(f)
	for s.Scan() {
		var i inst
		fs := strings.Fields(s.Text())
		i.c = aoc2023.Atoi(fs[1])
		switch fs[0] {
		case "U":
			i.dir = aoc2023.North
			curY -= i.c
			minY = min(minY, curY)
		case "R":
			i.dir = aoc2023.East
			curX += i.c
			maxX = max(maxX, curX)
		case "D":
			i.dir = aoc2023.South
			curY += i.c
			maxY = max(maxY, curY)
		case "L":
			i.dir = aoc2023.West
			curX -= i.c
			minX = min(minX, curX)
		default:
			panic(fmt.Errorf("unexpected first field: %q", s.Text()))
		}
		color := strings.TrimRight(strings.TrimLeft(fs[2], "(#"), ")")
		colorbyts, err := hex.DecodeString(color)
		if err != nil {
			panic(err)
		}
		i.color = [3]byte(colorbyts)
		insts = append(insts, i)
	}
	g := make(aoc2023.Grid2D, (maxY-minY)*2+1)
	for y := range g {
		g[y] = bytes.Repeat([]byte{'.'}, (maxX-minX*2)+1)
	}
	dig(g, insts)
	fmt.Println("Part 1:", count(g))
}

type inst struct {
	dir   aoc2023.Dir
	c     int
	color [3]byte
}

func dig(g aoc2023.Grid2D, insts []inst) {
	pos := aoc2023.Pos2D{X: len(g[0]) / 2, Y: len(g) / 2}
	for i, inst := range insts {
		var c byte
		switch inst.dir {
		case aoc2023.North, aoc2023.South:
			c = '|'
		case aoc2023.East, aoc2023.West:
			c = '-'
		}
		for j := 0; j < inst.c-1; j++ {
			pos = pos.Step(inst.dir)
			g.Set(pos, c)
		}
		pos = pos.Step(inst.dir)
		nextInst := insts[(i+1)%len(insts)]
		switch inst.dir {
		case aoc2023.North:
			switch nextInst.dir {
			case aoc2023.East:
				g.Set(pos, 'F')
			case aoc2023.West:
				g.Set(pos, '7')
			default:
				panic("impossible")
			}
		case aoc2023.South:
			switch nextInst.dir {
			case aoc2023.East:
				g.Set(pos, 'L')
			case aoc2023.West:
				g.Set(pos, 'J')
			default:
				panic("impossible")
			}
		case aoc2023.East:
			switch nextInst.dir {
			case aoc2023.North:
				g.Set(pos, 'J')
			case aoc2023.South:
				g.Set(pos, '7')
			default:
				panic("impossible")
			}
		case aoc2023.West:
			switch nextInst.dir {
			case aoc2023.North:
				g.Set(pos, 'L')
			case aoc2023.South:
				g.Set(pos, 'F')
			default:
				panic("impossible")
			}
		}
	}
	color(g)
}

func count(g aoc2023.Grid2D) int {
	var c int
	for _, row := range g {
		c += len(row) - bytes.Count(row, []byte{'O'})
	}
	return c
}

func color(g aoc2023.Grid2D) {
	for y := 0; y < len(g); y++ {
		var (
			inside bool
			start  byte // set when on top of a pipe
		)
		for x := 0; x < len(g[y]); x++ {
			p := aoc2023.Pos2D{X: x, Y: y}
			switch c, _ := g.Get(p); c {
			case '.':
				if start != 0 {
					panic("unexpected state")
				}
				if inside {
					g.Set(p, 'I')
				} else {
					g.Set(p, 'O')
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
