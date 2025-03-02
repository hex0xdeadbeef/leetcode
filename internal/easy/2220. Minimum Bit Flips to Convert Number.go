package easy

// https://leetcode.com/problems/minimum-bit-flips-to-convert-number/description/

func minBitFlips(start int, goal int) int {
	var (
		cnt int

		toBeAddedToMaxNum   int
		maxNum = max(start, goal)

		bitShift = 0
	)

	for toBeAddedToMaxNum != maxNum {
		if (start>>bitShift)&1 != (goal>>bitShift)&1 {
			cnt++
		}

		toBeAddedToMaxNum += maxNum & (1 << bitShift)
		bitShift++
	}

	return cnt
}

