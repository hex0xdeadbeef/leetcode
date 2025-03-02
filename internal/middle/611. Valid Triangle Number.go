package middle

// https://leetcode.com/problems/valid-triangle-number/description/

import (
	"sort"
)

func triangleNumber(nums []int) int {
	sort.Ints(nums)

	var (
		l, r int = 0, len(nums) - 1

		count int
	)

	for r-l >= 2 {
		validCount := isThereAppropriateNumber(nums[l], nums[r], nums[l+1:r])

		if validCount != -1 {
			count += validCount
		}

		l++

	}

	return count
}

func isThereAppropriateNumber(lNum, rNum int, subnums []int) int {
	var (
		l, r int = 0, len(subnums) - 1
		m    int

		previousValidIdx int
	)

	if l == r && isValidTriangle(lNum, rNum, subnums[m]) {
		return 1
	}

	for r-l > 0 {

		m = (l + r) / 2

		if isValidTriangle(lNum, rNum, subnums[m]) {
			previousValidIdx = m
			r = m
			continue
		}

		l = m
	}

	if isValidTriangle(lNum, rNum, subnums[l]) {
		return len(subnums) - previousValidIdx
	}

	return -1
}

func isValidTriangle(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}
