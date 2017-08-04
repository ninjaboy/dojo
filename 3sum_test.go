package main

import "testing"
import "sort"

func Test3Some(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(input)
	if len(res) != 2 {
		t.Error("Your solution sucked")
	}
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i1 := 0; (i1 < len(nums)-2) && nums[i1] <= 0; i1++ {
		if i1 > 0 && nums[i1] == nums[i1-1] {
			continue
		}
		i2 := i1 + 1
		i3 := len(nums) - 1
		target := -nums[i1]
		for i2 < i3 {
			if nums[i2]+nums[i3] == target {
				res = append(res, []int{nums[i1], nums[i2], nums[i3]})
				i2++
				i3--
			} else if nums[i2]+nums[i3] < target {
				i2++
			} else {
				i3--
			}

		}
	}
	return res
}
