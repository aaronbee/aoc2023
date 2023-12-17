package aoc2023

import (
	"cmp"
	"math/rand"
	"slices"
	"testing"
)

func TestHeap(t *testing.T) {
	h := NewHeap(cmp.Compare[int])

	vals := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Shuffle(len(vals), func(i, j int) {
		vals[i], vals[j] = vals[j], vals[i]
	})
	for _, v := range vals {
		h.Push(v)
	}
	var out []int
	for h.Len() > 0 {
		out = append(out, h.Pop())
	}
	if !slices.Equal(out, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}) {
		t.Errorf("unexpected output: %v", out)
	}
}
