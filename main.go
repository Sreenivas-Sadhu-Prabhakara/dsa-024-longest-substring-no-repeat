package main

import "fmt"

// LongestUniqueSubstring returns the length of the longest substring of s
// with no repeating character, using a sliding window that jumps its start
// past the previous occurrence of a repeated character.
// shortcut: byte-level window (ASCII); key the map by rune for full Unicode.
func LongestUniqueSubstring(s string) int {
	last := make(map[byte]int)
	best, start := 0, 0
	for i := 0; i < len(s); i++ {
		if j, ok := last[s[i]]; ok && j >= start {
			start = j + 1
		}
		last[s[i]] = i
		if i-start+1 > best {
			best = i - start + 1
		}
	}
	return best
}

func main() {
	fmt.Println(LongestUniqueSubstring("abcabcbb"))
}
