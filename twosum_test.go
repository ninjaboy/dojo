package main

import "testing"

func TestTwoSum(t *testing.T) {
	res := twoSum([]int{3, 2, 4}, 6)
	if res[0] != 1 || res[1] != 2 {
		t.Error("Wrong")
	}
}

func twoSum(nums []int, target int) []int {

	res := make(map[int]int, 0)
	for i, num := range nums {
		diff := target - num
		if val, ok := res[diff]; ok {
			return []int{val, i}
		}
		res[num] = i
	}
	return []int{}
}
