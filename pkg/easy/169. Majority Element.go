package easy

import (
	"slices"
)

// https://leetcode.com/problems/majority-element/description/

func majorityElement(nums []int) int {
	const (
		padding = 1
	)

	var (
		desirableMajorityCount = len(nums)/2 + padding

		curElem      = nums[0]
		curElemCount int
	)

	slices.Sort(nums)
	curElem = nums[len(nums)-1] + 1

	for _, v := range nums {

		if v != curElem {
			curElem = v
			curElemCount = 0
		}

		curElemCount++

		if curElemCount >= desirableMajorityCount {
			return curElem
		}
	}

	if curElemCount >= desirableMajorityCount {
		return curElem
	}

	panic("UNREACHABLE")
}
