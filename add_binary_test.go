package main

import "testing"

func TestAddBinary(t *testing.T) {
	a := "1"
	b := "1"

	c := addBinary(a, b)
	if c != "10" {
		t.Error("Incorrect result")
	}
}

func addBinary(a string, b string) string {
	res := ""
	sum, i, j := 0, len(a)-1, len(b)-1
	for sum == 1 || i >= 0 || j >= 0 {
		sum += getBitAt(a, i)
		i--
		sum += getBitAt(b, j)
		j--
		res = string(sum%2+int('0')) + res // 00 01 10 11
		sum /= 2
	}
	return res
}

func getBitAt(inputString string, index int) int {
	if index < len(inputString) && index >= 0 {
		return int(inputString[index] - '0')
	}
	return 0
}
