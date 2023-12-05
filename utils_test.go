package aoc2023

import (
	"slices"
	"testing"
)

func TestFieldsIter(t *testing.T) {
	var ss []string
	join := func(s string) { ss = append(ss, s) }
	FieldsIter("0", join)
	if !slices.Equal(ss, []string{"0"}) {
		t.Error(ss)
	}
	FieldsIter("  1 2 3  4 56 ", join)
	if !slices.Equal(ss, []string{"0", "1", "2", "3", "4", "56"}) {
		t.Error(ss)
	}
}
