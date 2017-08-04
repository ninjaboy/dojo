package main

import (
	"fmt"
	"testing"
)

func TestDotStarParsing(t *testing.T) {
	result := isMatch("ab", ".*c")
	fmt.Println(result)
}

func isMatch(s string, p string) bool {
	if s == "" && p == "" {
		return true
	}

	if len(p) > 1 && p[1] == '*' {
		return isMatch(s, p[2:]) ||
			(len(s) > 0 && (s[0] == p[0] || p[0] == '.') &&
				(isMatch(s[1:], p) ||
					isMatch(s[1:], p[2:])))

	}
	if len(p) > 0 && len(s) > 0 {
		if p[0] == '.' || s[0] == p[0] {
			return isMatch(s[1:], p[1:])
		}
	}

	return false
}
