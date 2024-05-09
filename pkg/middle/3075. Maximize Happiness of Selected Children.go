package middle

// https://leetcode.com/problems/maximize-happiness-of-selected-children/description/?envType=daily-question&envId=2024-05-09

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	const (
		idxPad = 1
	)

	var (
		sumOfhappiness int64
		r              int = len(happiness) - idxPad
	)

	sort.Ints(happiness)

	decreaseHappinesses(happiness[:r])

	for ; r >= 0 && k > 0; r-- {
		sumOfhappiness += int64(happiness[r])
		k--
	}

	return sumOfhappiness
}

func decreaseHappinesses(happinesses []int) {
	const (
		idxPad = 1

		lowBoundaryOfHappiness = 0
	)

	var curSubtrahend int = 1

	for i := len(happinesses) - idxPad; i >= 0; i-- {
		if happinesses[i]-curSubtrahend < lowBoundaryOfHappiness {
			happinesses[i] = lowBoundaryOfHappiness
			continue
		}

		happinesses[i] -= curSubtrahend

		curSubtrahend++
	}
}
