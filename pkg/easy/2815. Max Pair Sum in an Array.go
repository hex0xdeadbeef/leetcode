package easy

// https://leetcode.com/problems/max-pair-sum-in-an-array/description/

type NumbersPair struct {
	a, b int
}

func maxSum(nums []int) int {
	var (
		numsSortedByHighestDigit = make(map[int]NumbersPair, len(nums)/2)

		curDigit int
		np       NumbersPair

		maxSum int = -1
	)

	for _, v := range nums {
		curDigit = getMaxDigit(v)
		if _, ok := numsSortedByHighestDigit[curDigit]; !ok {
			numsSortedByHighestDigit[curDigit] = NumbersPair{a: v}
			continue
		}

		np = numsSortedByHighestDigit[curDigit]
		if np.a > np.b {
			np.b = max(np.b, v)
		} else {
			np.a = max(np.a, v)
		}
		numsSortedByHighestDigit[curDigit] = np
	}

	for _, v := range numsSortedByHighestDigit {
		if v.b == 0 {
			continue
		}
		maxSum = max(maxSum, v.a+v.b)
	}

	return maxSum
}

func getMaxDigit(n int) int {
	var (
		maxDigit int

		lastDigit int
	)

	for n > 0 {
		if lastDigit = n % 10; lastDigit > maxDigit {
			maxDigit = lastDigit
		}
		n /= 10
	}

	return maxDigit
}
