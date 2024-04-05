package middle

import "math"

// https://leetcode.com/problems/increasing-triplet-subsequence/submissions/1223237853/?envType=study-plan-v2&envId=leetcode-75

func increasingTriplet(nums []int) bool {
	triplet1, triplet2 := math.MaxInt32, math.MaxInt32

	for i := 0; i < len(nums); i++ {
		switch {
		case nums[i] <= triplet1:
			triplet1 = nums[i]
		case nums[i] <= triplet2:
			triplet2 = nums[i]
		default:
			return true
		}
	}
	return false
}
