package easy

import "math"

// https://leetcode.com/problems/maximum-average-subarray-i/?envType=study-plan-v2&envId=leetcode-75

func findMaxAverage(nums []int, k int) float64 {

	var (
		curSum int
		maxAvg float64
	)

	for i := 0; i < k; i++ {
		curSum += nums[i]
	}

	maxAvg = float64(curSum) / float64(k)

	for i := 1; i < len(nums)-1 && k+i-1 < len(nums); i++ {
		curSum -= nums[i-1]
		curSum += nums[k+i-1]

		if curAvg := float64(curSum) / float64(k); curAvg > maxAvg {
			maxAvg = curAvg
		}
	}

	return float64(maxAvg)
}

func findMaxAverageRepeat(nums []int, k int) float64 {
	const eps = 1e-9

	var (
		l, r, curSum int
		curAvg       float64
	)

	maxAvg := -math.MaxFloat64

	if len(nums) == 1 {
		return float64(nums[0])
	}

	for ; r < k && r < len(nums); r, curSum = r+1, curSum+nums[r] {
	}

	maxAvg = float64(curSum) / float64(k)

	for r < len(nums) {
		curSum = curSum - nums[l] + nums[r]

		curAvg = float64(curSum) / float64(k)

		if curAvg > maxAvg+eps {
			maxAvg = curAvg
		}

		l, r = l+1, r+1
	}

	return roundTo5(maxAvg)
}

func roundTo5(x float64) float64 {
	return math.Round(x*1e5) / 1e5
}
