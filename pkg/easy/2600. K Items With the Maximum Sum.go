package easy

// https://leetcode.com/problems/k-items-with-the-maximum-sum/description/

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	var (
		sum int
	)

	if k == 0 {
		return 0
	}

	if numOnes != 0 {
		onesCnt := min(numOnes, k)
		sum += onesCnt
		k -= onesCnt
	}

	if k > 0 {
		k -= min(numZeros, k)
	}

	if k > 0 && numNegOnes != 0 {
		sum += -1 * k
	}

	return sum
}
