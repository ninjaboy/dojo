package main

import "testing"

func TestVerboseNum(t *testing.T) {
	i := 123
	res := numberToWords(i)
	if res != "One Hundred Twenty Three" {
		t.Error("YOu suck")
	}
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	nums_large := []string{"", "Thousand", "Million", "Billion"}
	res := ""
	for i := 0; num > 0; i++ {
		var left string
		if num%1000 > 0 {
			left = appendString(verbose(num%1000), nums_large[i])
		}

		res = appendString(left, res)
		num /= 1000
	}
	return res
}

func verbose(num int) string {
	nums_1_19 := []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	nums_20_90 := []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	res := ""
	if num/100 > 0 {
		res = appendString(nums_1_19[num/100], "Hundred")
	}
	num %= 100
	if num < 20 {
		res = appendString(res, nums_1_19[num%100])
	} else {
		res = appendString(appendString(res, nums_20_90[num/10]), nums_1_19[num%10])
	}
	//res = strings.TrimSpace(res)
	return res
}

func appendString(left string, right string) string {
	if left == "" && right == "" {
		return ""
	}
	if left == "" {
		return right
	}
	if right == "" {
		return left
	}
	return left + " " + right
}

// 1..19 20..99
// 100..999
