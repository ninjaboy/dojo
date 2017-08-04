package main

import "testing"

var mapping map[byte]string

func TestCombinations(t *testing.T) {
	combs := letterCombinations("2")
	if len(combs) != 3 {
		t.Error("Failed")
	}
}

func letterCombinations(digits string) []string {
	mapping = initMap()
	combs := make(chan string)

	go getCombinationsAsync(digits, combs, 0, "")
	var res []string
	for item := range combs {
		res = append(res, item)
	}
	return res
}

func getCombinationsAsync(digits string, combs chan string, level int, current string) {
	getCombinations(digits, combs, 0, "")
	close(combs)
}

func func1(c []string) {
	c = append(c, "test")
}

func getCombinations(digits string, combs chan string, level int, current string) {
	if level <= len(digits)-1 {
		for _, c := range mapping[byte(digits[level])] {
			if level == len(digits)-1 {
				combs <- current + string(c)
			} else {
				getCombinations(digits, combs, level+1, current+string(c))
			}
		}
	}
}

func initMap() map[byte]string {
	mapping := make(map[byte]string, 10)
	mapping['0'] = " "
	mapping['1'] = ""
	mapping['2'] = "abc"
	mapping['3'] = "def"
	mapping['4'] = "ghi"
	mapping['5'] = "jkl"
	mapping['6'] = "mno"
	mapping['7'] = "prs"
	mapping['8'] = "tuv"
	mapping['9'] = "wxyz"
	return mapping
}
