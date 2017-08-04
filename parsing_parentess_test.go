package main

import "testing"

func TestParentess(t *testing.T) {

	checkInput(t, "(()", []string{"()"})
	checkInput(t, "(((a))", []string{"((a))"})
	checkInput(t, ")(((a))", []string{"((a))"})
	checkInput(t, "(()()))", []string{"(()())", "((()))"})
	checkInput(t, "(a()()))", []string{"(a()())", "(a(()))"})

}

func checkInput(t *testing.T, input string, expected []string) {
	res := fixExpression(input)
	for _, v := range expected {
		if _, ok := res[v]; !ok {
			t.Errorf("Expected item: %v not found", v)
		}
	}
}

func fixExpression(input string) map[string]bool {
	res := make(map[string]bool, 0)
	l, r := calculatePars(input)
	check(input, 0, res, l, r, 0)

	// resArr := getArr(res)
	return res
}

func getArr(m map[string]bool) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func check(input string, i int, res map[string]bool, l int, r int, c int) {
	if i == len(input) {
		if c == 0 {
			if _, ok := res[input]; !ok {
				res[input] = true
			}
		}
		return
	}

	switch {
	case input[i] == '(':
		if l > 0 { //can try remove l and check
			check(inputRm(input, i), i, res, l-1, r, c)
		}
		check(input, i+1, res, l, r, c+1)
	case input[i] == ')':
		if r > 0 { //can try to remove r and check
			check(inputRm(input, i), i, res, l, r-1, c)
		}
		if c > 0 { //if count is less than or equal to zero means we are in negative situation already. so skip
			check(input, i+1, res, l, r, c-1)
		}
	default:
		check(input, i+1, res, l, r, c)
	}
}

func inputRm(input string, i int) string {
	res := input[0:i]
	if i+1 < len(input) {
		res += input[i+1 : len(input)]
	}
	return res
}

func calculatePars(input string) (l int, r int) {
	for _, c := range input {
		switch {
		case c == '(':
			l++
		case c == ')':
			if l != 0 {
				l--
			} else {
				r++
			}
		}
	}
	return
}
