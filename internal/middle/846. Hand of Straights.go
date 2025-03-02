package middle

// https://leetcode.com/problems/hand-of-straights/description/?envType=daily-question&envId=2024-06-06

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	var (
		cnts = make(map[int]int, len(hand))

		stack     = make([]int, 0, len(hand))
		curMinCnt int
	)

	for _, v := range hand {
		cnts[v]++
	}

	sort.Ints(hand)
	hand = filterSameNums(hand)

	for i := len(hand) - 1; i >= 0; i-- {
		if len(stack) == 0 {
			stack = append(stack, hand[i])
			curMinCnt = cnts[hand[i]]
			continue
		}

		if len(stack) == groupSize {
			stack = processStack(stack, cnts, groupSize, curMinCnt)
			for _, v := range stack {

				if curMinCnt = min(curMinCnt, cnts[v]); curMinCnt == 0 {
					return false
				}

			}

			i++
			continue
		}

		if stackTop := stack[len(stack)-1]; stackTop-1 == hand[i] {
			stack, curMinCnt = append(stack, hand[i]), min(curMinCnt, cnts[hand[i]])
		}

	}

	stack = processStack(stack, cnts, groupSize, curMinCnt)
	if len(stack) != 0 {
		return false
	}

	return true
}

func processStack(stack []int, cnts map[int]int, groupSize, minCnt int) []int {
	if len(stack) != groupSize {
		return stack
	}

	for _, v := range stack {
		cnts[v] -= minCnt

		if cnts[v] == 0 {
			stack = stack[1:]
		}
	}

	return stack
}

func filterSameNums(nums []int) []int {
	var (
		filtered = make([]int, 0, len(nums))
		cur      = nums[0]
	)
	filtered = append(filtered, cur)

	for i := 0; i < len(nums); i++ {
		if nums[i] == cur {
			continue
		}

		cur = nums[i]
		filtered = append(filtered, cur)
	}

	return filtered
}

func decrementCnts(l, r int, cnts map[int]int) {
	for i := l; i <= r; i++ {
		cnts[i]--
	}
}
