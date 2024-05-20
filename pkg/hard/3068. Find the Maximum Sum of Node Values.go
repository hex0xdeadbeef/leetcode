package hard

import (
	"slices"
)

type PairXOR struct {
	before, after int

	difference int
}

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	var (
		pairs   = make([]PairXOR, len(nums))
		newPair PairXOR

		idx, firstPosIdx int

		lastNegPair, pairAfterNegatives PairXOR
		resSum                          int64
	)

	for i, v := range nums {
		newPair = PairXOR{before: v, after: v ^ k, difference: v ^ k - v}
		pairs[i] = newPair
	}

	slices.SortFunc(pairs, func(a, b PairXOR) int {
		return a.difference - b.difference
	})

	for idx = 0; idx < len(pairs) && pairs[idx].difference < 0; idx++ {
		if idx+1 < len(pairs) && pairs[idx+1].difference > 0 {
			lastNegPair, pairAfterNegatives = pairs[idx], pairs[idx+1]
			firstPosIdx = idx + 1
		}
	}

	if (len(pairs)-idx)%2 == 1 {
		idx++
	}

	for i := 0; i < idx; i++ {
		resSum += int64(pairs[i].before)
	}

	for ; idx < len(pairs); idx++ {
		resSum += int64(pairs[idx].after)
	}

	if (len(pairs)-firstPosIdx)%2 == 1 {
		postSum := (resSum - (int64(lastNegPair.before) + int64(pairAfterNegatives.before)) + int64(lastNegPair.after) + int64(pairAfterNegatives.after))

		if resSum < postSum {
			resSum = postSum
		}

	}

	return int64(resSum)
}
