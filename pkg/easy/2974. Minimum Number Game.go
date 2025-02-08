package easy

// https://leetcode.com/problems/minimum-number-game/submissions/1500757168/

import (
	"slices"
)

func numberGame(nums []int) []int {
	var res []int

	slices.Sort(nums)

	for i := 0; i < len(nums); i += 2 {
		if i+1 < len(nums) {
			res = append(res, nums[i+1])
		}

		res = append(res, nums[i])
	}

	return res
}
