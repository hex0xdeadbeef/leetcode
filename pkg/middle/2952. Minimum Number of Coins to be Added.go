package middle

import "sort"

func minimumAddedCoins(coins []int, target int) int {
	const (
		lengthDivider = 2

		pad = 1
	)

	var (
		added = make([]int, 0, len(coins)/lengthDivider)
	)

	sort.Ints(coins)

	var (
		curTarget int = 0 // The number we want to get [1, n]
	)
	for _, v := range coins {
		for curTarget+pad < v {
			curTarget += curTarget + pad
			added = append(added, curTarget)
			if curTarget >= target {
				break
			}
		}

		if curTarget+pad == v {
			curTarget += v
		}

		if curTarget >= target {
			break
		}
	}

	for curTarget < target {
		curTarget += curTarget + pad
		added = append(added, curTarget)
	}

	return len(added)
}
