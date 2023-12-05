package aoc2023

import (
	"slices"
	"testing"
)

func TestIntFieldsIter(t *testing.T) {
	var ss []int
	join := func(v int) { ss = append(ss, v) }
	IntFieldsIter("0", join)
	if !slices.Equal(ss, []int{0}) {
		t.Error(ss)
	}
	IntFieldsIter("  1 2 3  4 56 ", join)
	if !slices.Equal(ss, []int{0, 1, 2, 3, 4, 56}) {
		t.Error(ss)
	}
}
