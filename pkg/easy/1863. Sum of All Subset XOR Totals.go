package easy

// https://leetcode.com/problems/sum-of-all-subset-xor-totals/description/?envType=daily-question&envId=2024-05-20

import "math"

func subsetXORSum(nums []int) int {
	var (
		curNum, idx, curXOR int

		resSum int
	)

	for i := 0; i < int(math.Pow(2, float64(len(nums)))); i++ {
		idx, curNum, curXOR = 0, i, 0

		for curNum != 0 {
			if curNum&1 == 1 {
				curXOR ^= nums[idx]
			}
			idx++

			curNum >>= 1
		}

		resSum += curXOR

	}

	return resSum
}
