package easy

// https://leetcode.com/problems/smallest-missing-integer-greater-than-sequential-prefix-sum/

// type empty struct{}

func missingInteger(nums []int) int {

	const (
		idxPad = 1

		one = 1
	)

	if len(nums) == one {
		return nums[0] + one
	}

	var (
		cur         = nums[0] - one
		sumOfPrefix int

		rightSideOfPrefix map[int]empty
		startIdx          int
	)

	for i, v := range nums {
		if v-cur == one {
			sumOfPrefix += v
			cur = v
			startIdx = i
			continue
		}

		startIdx++
		break
	}

	rightSideOfPrefix = make(map[int]empty, len(nums)-startIdx)
	for i := 0; i < len(nums); i++ {
		rightSideOfPrefix[nums[i]] = empty{}
	}

	var curTarget int = sumOfPrefix
	for range rightSideOfPrefix {
		if _, ok := rightSideOfPrefix[curTarget]; ok {
			curTarget++
			continue
		}

		return curTarget
	}

	return curTarget
}
