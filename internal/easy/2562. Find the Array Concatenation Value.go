package easy

// https://leetcode.com/problems/find-the-array-concatenation-value/

import "strconv"

func findTheArrayConcVal(nums []int) int64 {

	var (
		l, r = 0, len(nums) - 1
		sum  int64

		leftNumStr, rightNumStr string
		curConcatSum            int64
	)

	for r-l >= 1 {
		leftNumStr, rightNumStr = strconv.Itoa(nums[l]), strconv.Itoa(nums[r])
		curConcatSum, _ = strconv.ParseInt(leftNumStr+rightNumStr, 10, 64)
		sum += curConcatSum
		r--
		l++
	}

	if r-l == 0 {
		sum += int64(nums[l])
	}

	return sum
}
