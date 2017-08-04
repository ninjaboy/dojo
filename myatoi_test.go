package main

import "testing"

func TestMyAtoi(t *testing.T) {
	//input := `214748364`
	// res := myAtoi(input)
	// if res != 0 {
	// 	t.Error("Baaad baaad motherfucker")
	// }
}

const MaxInt32 = int(int32(^uint32(0) >> 1))
const MinInt32 = -MaxInt32 - 1

func myAtoi(str string) int {
	var res int
	sign := 1
	if len(str) == 0 {
		return 0
	}
	str = trim(str)
	if str[0] == '-' {
		str = str[1:]
		sign = -1
	} else if str[0] == '+' {
		str = str[1:]
	}
	for _, c := range str {
		if !isValid(c) {
			break
		}
		i := toInt(c)
		if sign > 0 && (res*10+i > MaxInt32) {
			res = MaxInt32
		} else if sign < 0 && (res*10+i > -MinInt32) {
			res = -MinInt32
		} else {
			res *= 10
			res += i
		}
	}

	res = res * sign
	return res

}

func trim(s string) string {
	skip := 0
	for _, c := range s {
		if c == ' ' || c == '\t' {
			skip++
		} else {
			break
		}
	}
	return s[skip:]
}

func toInt(c rune) int {
	return int(c - '0')
}

func isValid(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}
