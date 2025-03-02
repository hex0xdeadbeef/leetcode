package middle

import . "leetcode/pkg/datastructures/doublylinkedlist"

func longestOnes(nums []int, k int) int {
	var l, r, maxLen int
	list := New()

	for ; r < len(nums); r++ {
		if nums[r] == 1 {
			continue
		}

		if k == 0 {
			break
		}

		list.InsertTail(r)
		k--

	}
	maxLen = r - l

	for ; r < len(nums); r++ {
		if nums[r] == 1 {
			continue
		}

		maxLen = max(maxLen, r-l-1)

		if list.Count() == 0 {
			l = r
			continue
		}

		list.InsertTail(r)
		l = list.PopHead().Val

	}

	return max(maxLen, r-l-1)
}
