package middle

// https://leetcode.com/problems/subsets/description/

import (
	"math"
)

func subsets(nums []int) [][]int {

	var (
		curNum int

		curSubset []int
		idx       int

		res [][]int
	)

	for i := 0; i < int(math.Pow(2, float64(len(nums)))); i++ {
		idx, curNum, curSubset = 0, i, make([]int, 0, len(nums)/2)

		for curNum != 0 {
			if curNum&1 == 1 {
				curSubset = append(curSubset, nums[idx])
			}
			idx++

			curNum >>= 1
		}

		res = append(res, curSubset)

	}

	return res
}
