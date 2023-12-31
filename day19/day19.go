package main

import (
	"bufio"
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
	ws := workflows{
		m: make(map[string]int),
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		if len(s.Bytes()) == 0 {
			break
		}
		w := parseWorkflow(s.Text())
		ws.m[w.name] = len(ws.l)
		ws.l = append(ws.l, w)
	}
	var part1 int
	for s.Scan() {
		p := parsePart(s.Text())
		if ws.run(p) {
			part1 += p.sum()
		}
	}
	fmt.Println("Part 1:", part1)
}

type workflows struct {
	m map[string]int
	l []*workflow
}

type workflow struct {
	name string
	rs   []rule
}

type rule struct {
	category byte
	op       byte
	val      int
	result   string // "A|R" or workflow name
}

type part struct {
	x, m, a, s int
}

func (p *part) sum() int {
	return p.x + p.m + p.a + p.s
}

func mustCut(l, d string) (string, string) {
	left, right, ok := strings.Cut(l, d)
	if !ok {
		panic(fmt.Errorf("expected to find %q in %q", d, l))
	}
	return left, right
}

func parseWorkflow(l string) *workflow {
	name, rest := mustCut(l, "{")
	w := workflow{name: name}
	for _, r := range strings.Split(rest, ",") {
		if s, ok := strings.CutSuffix(r, "}"); ok {
			w.rs = append(w.rs, rule{result: s})
			break
		}
		cond, result := mustCut(r, ":")
		w.rs = append(w.rs, rule{
			category: cond[0],
			op:       cond[1],
			val:      aoc2023.Atoi(cond[2:]),
			result:   result,
		})
	}

	return &w
}

func parsePart(l string) *part {
	x, rest := mustCut(strings.TrimPrefix(l, "{x="), ",")
	m, rest := mustCut(strings.TrimPrefix(rest, "m="), ",")
	a, rest := mustCut(strings.TrimPrefix(rest, "a="), ",")
	s, _ := mustCut(strings.TrimPrefix(rest, "s="), "}")
	return &part{x: aoc2023.Atoi(x), m: aoc2023.Atoi(m), a: aoc2023.Atoi(a), s: aoc2023.Atoi(s)}
}

// run returns "A"=accept, "R"=reject, ""=continue, or next rule
func (r rule) run(p *part) string {
	if r.category == 0 {
		return r.result
	}
	var v int
	switch r.category {
	case 'x':
		v = p.x
	case 'm':
		v = p.m
	case 'a':
		v = p.a
	case 's':
		v = p.s
	default:
		panic("unknown category")
	}
	switch r.op {
	case '<':
		if v < r.val {
			return r.result
		}
	case '>':
		if v > r.val {
			return r.result
		}
	}
	return ""
}

func (w *workflow) run(p *part) string {
	for _, r := range w.rs {
		if res := r.run(p); res != "" {
			return res
		}
	}
	panic("workflow reached the end")
}

func (ws *workflows) run(p *part) bool {
	w := ws.l[ws.m["in"]]
	for {
		switch res := w.run(p); res {
		case "A":
			return true
		case "R":
			return false
		default:
			w = ws.l[ws.m[res]]
		}
	}
}
