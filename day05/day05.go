package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"

	"github.com/aaronbee/aoc2023"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	if !s.Scan() {
		panic(fmt.Errorf("unexpected scan end: %s", s.Err()))
	}
	var seeds []seed
	aoc2023.IntFieldsIter(strings.TrimPrefix(s.Text(), "seeds: "), func(i int) {
		seeds = append(seeds, seed{s: i, c: 1})
	})

	var ms []*maps
	var cur *maps
	for s.Scan() {
		l := s.Text()
		if l == "" {
			if cur != nil {
				slices.SortFunc(cur.ms, func(a, b *mapping) int {
					return cmp.Compare(a.src, b.src)
				})
			}
			continue
		}
		if name, ok := strings.CutSuffix(l, " map:"); ok {
			cur = &maps{name: name}
			ms = append(ms, cur)
			continue
		}
		var m mapping
		fmt.Sscanf(l, "%d %d %d", &m.dst, &m.src, &m.count)
		cur.ms = append(cur.ms, &m)
	}
	slices.SortFunc(cur.ms, func(a, b *mapping) int {
		return cmp.Compare(a.src, b.src)
	})

	fmt.Println("Part 1:", closestLocation(seeds, ms))
	seeds2 := make([]seed, len(seeds)/2)
	for i := 0; i < len(seeds); i += 2 {
		s := seeds[i]
		c := seeds[i+1]
		seeds2[i/2] = seed{s.s, c.s}
	}
	fmt.Println("Part 2:", closestLocation(seeds2, ms))
}

type seed struct {
	s int
	c int
}

type maps struct {
	name string
	ms   []*mapping
}

type mapping struct {
	dst   int
	src   int
	count int
}

func closestLocation(seeds []seed, ms []*maps) int {
	closestLocation := math.MaxInt
	for _, seed := range seeds {
		closestLocation = min(closestLocation, seedLocation(seed, ms))
	}
	return closestLocation
}

func seedLocation(s seed, ms []*maps) int {
	if len(ms) == 0 {
		return s.s
	}
	closestLocation := math.MaxInt
	maps, nextMs := ms[0], ms[1:]

	for _, m := range maps.ms {
		if s.s > m.dst+m.count {
			continue
		}
		if s.s < m.src {
			// Some seeds before mapping
			if s.s+s.c-1 < m.src {
				// All seeds before mapping
				return min(closestLocation, seedLocation(s, nextMs))
			}
			nextS := m.src
			curSeed := seed{s: s.s, c: nextS - s.s}
			closestLocation = min(closestLocation, seedLocation(curSeed, nextMs))
			s = seed{s: nextS, c: s.s + s.c - nextS}
		}
		mapped := s.s + m.dst - m.src
		if s.s+s.c < m.src+m.count {
			// Rest of seeds contained in this mapping
			s = seed{s: mapped, c: s.c}
			return min(closestLocation, seedLocation(s, nextMs))
		}
		nextS := m.src + m.count
		curSeed := seed{s: mapped, c: nextS - s.s}
		closestLocation = min(closestLocation, seedLocation(curSeed, nextMs))
		s = seed{s: nextS, c: s.s + s.c - nextS}
	}
	// If we've reached here, there are some seeds beyond any mapping
	return min(closestLocation, seedLocation(s, nextMs))
}
