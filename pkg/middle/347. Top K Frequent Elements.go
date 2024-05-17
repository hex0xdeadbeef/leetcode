package middle

// https://leetcode.com/problems/top-k-frequent-elements/description/

import "sort"

type FreqPair struct {
	val int
	cnt int
}

type FreqPairs []*FreqPair

func (a FreqPairs) Len() int           { return len(a) }
func (a FreqPairs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a FreqPairs) Less(i, j int) bool { return a[i].cnt > a[j].cnt }

func topKFrequent(nums []int, k int) []int {
	var (
		seen = make(map[int]*FreqPair, len(nums))

		res []*FreqPair
	)

	for _, v := range nums {
		if pair, ok := seen[v]; ok {
			pair.cnt++
			continue
		}

		seen[v] = &FreqPair{val: v, cnt: 1}
	}

	var (
		idx int
	)
	res = make([]*FreqPair, len(seen))
	for _, v := range seen {
		res[idx] = v
		idx++
	}
	res = res[:idx]

	sort.Sort(FreqPairs(res))
	nums = nums[:0]
	for i := 0; k > 0; i++ {
		nums = append(nums, res[i].val)
		k--
	}

	return nums
}
