package easy

// https://leetcode.com/problems/minimum-bit-flips-to-convert-number/description/

func minBitFlips(start int, goal int) int {
	var (
		cnt int

		target   int
		addition = max(start, goal)

		bitShift = 0
	)

	for target != addition {
		if (start>>bitShift)&1 != (goal>>bitShift)&1 {
			cnt++
		}

		target += addition & (1 << bitShift)
		bitShift++
	}

	return cnt
}
