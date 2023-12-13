package main

import "testing"

func TestMatchCount(t *testing.T) {
	for _, tc := range []struct {
		s   string
		cs  []int
		exp int
	}{{
		s: "????.?#???#??", cs: []int{1, 2, 6},
		exp: 2,
	}, {
		s: "????#?#?..?#?", cs: []int{4, 1},
		exp: 2,
	}} {
		got := matchCount(tc.s, tc.cs)
		if got != tc.exp {
			t.Errorf("%s %v got=%d exp=%d", tc.s, tc.cs, got, tc.exp)
		}
	}
}
