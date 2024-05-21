package middle

// https://leetcode.com/problems/3sum/description/

import (
	"slices"
)

// type empty struct{}

func threeSum(nums []int) [][]int {

	const (
		target = 0
	)

	slices.SortFunc(nums, func(a, b int) int {
		return a - b
	})

	var (
		curApproximation int
		l, r             int

		triplet [3]int
		seen    = make(map[[3]int]empty, len(nums)/3)

		preRes [][3]int = make([][3]int, 0, len(nums)/3)
	)

	for i := 0; i < len(nums)-2; i++ {
		l, r = i+1, len(nums)-1

		for l < r {
			curApproximation = 0 + (nums[i] + nums[l] + nums[r])

			if curApproximation < 0 {
				l++
				continue
			}

			if curApproximation > 0 {
				r--
				continue
			}

			triplet = [3]int{nums[i], nums[l], nums[r]}
			if _, ok := seen[triplet]; !ok {
				preRes = append(preRes, triplet)
				seen[triplet] = empty{}
			}

			l++
			r--
		}
	}

	var res [][]int = make([][]int, len(preRes))
	for i, v := range preRes {
		res[i] = v[:]
	}

	return res
}
