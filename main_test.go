package main

import "testing"

func TestLongestUniqueSubstring(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"a", 1},
		{"abba", 2},
		{"dvdf", 3},
	}
	for _, c := range cases {
		if got := LongestUniqueSubstring(c.in); got != c.want {
			t.Errorf("LongestUniqueSubstring(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}
