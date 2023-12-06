package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
	"time"

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
		seeds = append(seeds, seed{i: i, c: 1})
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
		seeds2[i/2] = seed{s.i, c.i}
	}
	fmt.Println("Part 1:", closestLocation(seeds2, ms))
}

type seed struct {
	i int
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
		fmt.Println(time.Now(), seed)
		for i := 0; i < seed.c; i++ {
			closestLocation = min(closestLocation, seedLocation(seed.i+i, ms))
		}
	}

	return closestLocation
}

func seedLocation(seed int, ms []*maps) int {
	v := seed
	for _, ms := range ms {
		i, ok := slices.BinarySearchFunc(ms.ms, v, func(m *mapping, v int) int {
			if m.src > v {
				return +1
			} else if m.src+m.count <= v {
				return -1
			}
			return 0
		})
		if !ok {
			// no mapping found
			continue
		}
		m := ms.ms[i]
		v += m.dst - m.src
	}
	return v
}
