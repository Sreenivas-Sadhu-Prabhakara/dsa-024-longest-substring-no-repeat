# Longest Substring Without Repeating Characters

> Striver SDE Sheet 024 · Day 4 Arrays — sliding window whose start jumps straight past the last occurrence of the repeated character. Go, stdlib only, tested.

## Problem

Find the length of the longest substring of a string that contains no repeated character.

## Approach

Slide a window [start, i]. On each new character, if it was last seen *inside* the current window, the window start jumps to just past that occurrence — no character-by-character shrinking. The `j >= start` guard is the subtle part: an occurrence *before* the window (as in `abba` when the second `a` arrives after start already moved past the first `b`) must not drag the start backwards. One map from character to last index does all the bookkeeping.

**Complexity:** O(n) time, O(min(n, alphabet)) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
