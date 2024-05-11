package easy

// https://leetcode.com/problems/sum-of-digits-in-base-k/description/

import "strconv"

func SumBase(n int, k int) int {

	const (
		zeroByteVal int = 48
	)

	var (
		nStrInBase = strconv.FormatInt(int64(n), k)

		sum int
	)

	for i := 0; i < len(nStrInBase); i++ {
		sum += int(nStrInBase[i]) - zeroByteVal
	}

	return sum
}
