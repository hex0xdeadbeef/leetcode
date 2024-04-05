package easy

// https://leetcode.com/problems/max-number-of-k-sum-pairs/description/?envType=study-plan-v2&envId=leetcode-75

func FindMaxAverage(nums []int, k int) float64 {

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
