package middle

func minSubArrayLen(target int, nums []int) int {
	var l, r, curSum, minLen int

	for ; r < len(nums) && curSum < target; r++ {
		curSum += nums[r]
	}

	if minLen = r - l; minLen == len(nums) && curSum < target {
		return 0
	}

	for l < r {
		curSum, l = curSum-nums[l], l+1
		if curSum >= target && r-l < minLen {
			minLen = r - l
		}

		for ; r < len(nums) && curSum < target; r++ {
			curSum += nums[r]
		}
		if curSum >= target && r-l < minLen {
			minLen = r - l
		}

	}

	return minLen
}
