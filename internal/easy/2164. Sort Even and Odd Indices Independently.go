package easy

// https://leetcode.com/problems/sort-even-and-odd-indices-independently/description/

import (
	"sort"
)

type Pair struct {
	val int
	idx int
}

type Pairs []Pair

func (a Pairs) Len() int {
	return len(a)
}

func (a Pairs) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Pairs) Less(i, j int) bool {
	if a[i].idx%2 == 0 {
		return a[i].val < a[j].val
	}
	return a[i].val > a[j].val
}

func sortEvenOdd(nums []int) []int {

	var (
		odds, evens = make([]Pair, 0, len(nums)/2), make([]Pair, 0, len(nums)/2)
		newPair     Pair
	)

	for i, v := range nums {
		newPair = Pair{val: v, idx: i}

		if i%2 == 1 {
			odds = append(odds, newPair)
			continue
		}

		evens = append(evens, newPair)
	}

	sort.Sort(Pairs(odds))
	sort.Sort(Pairs(evens))

	var (
		oP, eP int
	)

	nums = nums[:0]

	for i := 0; i < len(odds)+len(evens); i++ {
		switch i % 2 {
		case 0:
			nums = append(nums, evens[eP].val)
			eP++
		case 1:
			nums = append(nums, odds[oP].val)
			oP++
		}
	}

	return nums
}
