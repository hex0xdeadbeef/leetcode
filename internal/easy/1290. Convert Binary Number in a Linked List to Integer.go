package easy

import (
	. "leetcode/pkg/datastructures/linkedlist"
	"math"
)

func getDecimalValue(head *ListNode) int {
	head = reverseList(head)

	var (
		res    int
		curPow int
	)

	for ; head != nil; head = head.Next {
		res += head.Val * int(math.Pow(2, float64(curPow)))
		curPow++
	}

	return res
}
