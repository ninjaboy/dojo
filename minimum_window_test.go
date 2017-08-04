package main

import "strings"
import "testing"

type match struct {
	v int
	s int
	e int
}

func TestMinWindows(t *testing.T) {
	if res := minWindow("1", "1"); res != "1" {
		t.Error("You sucked")
	}
}

func minWindow(s string, t string) string {
	var best *match
	var results []match
	chksum := (1 << uint(len(t))) - 1
	for i, c := range s {
		if ci := strings.IndexRune(t, c); ci != -1 {
			results = append(results, match{v: 0, s: i, e: 0})
			for _, r := range results {
				r.v |= (1 << uint(ci))
				if r.v == chksum {
					r.e = i + 1
					if best == nil || best.e-best.s > r.e-r.s {
						best = &r
					}
				}
			}
		}
	}
	if best == nil {
		return ""
	}
	return s[best.s:best.e]
}
