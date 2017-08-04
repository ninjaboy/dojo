package main

import (
	"fmt"
	"testing"
)

func TestDecodeWays(t *testing.T) {
	// res := numDecodings("9371597631128776948387197132267188677349946742344217846154932859125134924241649584251978418763151253")
	// if res != 2 {
	// 	t.Error("You suck")
	// }
}

func numDecodings(s string) int {
	if s == "" {
		return 0
	}
	res := make(chan string)
	go getDecodingsAsync(s, res)
	i := 0
	for r := range res {
		fmt.Println(r)
		i++
	}
	return i
}

func getDecodingsAsync(s string, res chan string) {
	getDecodings(s, res, 0, "", false)
	close(res)
}

func getDecodings(s string, res chan string, cur int, curRes string, mid bool) {
	if cur == len(s) {
		if !mid {
			res <- curRes
		}
		return
	}
	if mid {
		i := stringToInt(s[cur-1 : cur+1])
		if i >= 10 && i <= 26 {
			curRes += intDecode(i)
			getDecodings(s, res, cur+1, curRes, false)
		}
	} else {
		if s[cur] != '0' {
			getDecodings(s, res, cur+1, curRes+intDecode(stringToInt(s[cur:cur+1])), false)
		}
		if s[cur] >= '0' && s[cur] <= '1' {
			getDecodings(s, res, cur+1, curRes, true)
		}
	}
}

func intDecode(i int) string {
	return string('A' + i - 1)
}

func stringToInt(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res *= 10
		res += digitCharToInt(s[i])
	}
	return res
}

func digitCharToInt(c byte) int {
	return int(c - '0')
}
