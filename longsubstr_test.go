package main

import "testing"

func TestLongestSubstr(t *testing.T) {
	data := "somedata"
	data1 := data[0:0]
	data2 := data[0:1]
	if data1 == data2 {
		return
	}
	max := lengthOfLongestSubstring("tmmzuxt")
	if max != 5 {
		t.Error("Expected length for 'tmmzuxt' is 5")
	}
}

func lengthOfLongestSubstring(s string) int {
	positions := make(map[rune]int)
	start := 0
	max := 0
	for i, c := range s {
		if _, ok := positions[c]; ok && start < positions[c] {
			start = positions[c] + 1
		} else {
			max = maxval(max, i-start+1)
		}
		positions[c] = i
	}
	return max
}

func maxval(i, j int) int {
	if i >= j {
		return i
	} else {
		return j
	}
}
