package easy

// https://leetcode.com/problems/maximum-count-of-positive-integer-and-negative-integer/description/

func maximumCount(nums []int) int {
	var (
		posCnt, negCnt int
		l, r           = 0, len(nums) - 1
	)

	for {
		if nums[l] < 0 {
			l++
			negCnt++
		}

		if nums[r] > 0 {
			posCnt++
			r--
		}

		if (l == 0 && r == -1) || (l == len(nums) && r == len(nums)-1) {
			break
		}

		if v1, v2 := nums[l], nums[r]; (v1 > 0 || v1 == 0) && (v2 < 0 || v2 == 0) {
			break
		}

	}

	return max(posCnt, negCnt)
}
