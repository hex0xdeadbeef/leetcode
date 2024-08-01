package easy

// https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-i/description/

func maximumTripletValue(nums []int) int64 {
	var (
		maxs = getMaxs(nums)
		lMax = nums[0]

		maxProduct int64 = -1e9
	)

	// a = lMax, b = nums[i], c = maxs[i+1]
	for i := 1; i < len(nums)-1; i++ {
		if curProd := int64(lMax-nums[i]) * int64(maxs[i+1]); curProd > maxProduct {
			maxProduct = curProd
		}

		lMax = max(lMax, nums[i])
	}

	if maxProduct < 0 {
		return 0
	}

	return maxProduct
}

func getMaxs(nums []int) []int {
	var (
		copiedNums = append([]int(nil), nums...)
	)

	for i := len(copiedNums) - 2; i > 0; i-- {
		copiedNums[i] = max(copiedNums[i+1], copiedNums[i])
	}

	return copiedNums
}
