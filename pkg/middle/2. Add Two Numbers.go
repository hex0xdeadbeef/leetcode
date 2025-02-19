package middle

// https://leetcode.com/problems/add-two-numbers/description/

import (
	. "leetcode/pkg/datastructures/linkedlist"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		draftNum int

		res    = &ListNode{Val: 0}
		preRes *ListNode

		head = res
	)

	for l1 != nil && l2 != nil {
		var (
			nodesSum int
		)

		if l1 != nil {
			nodesSum += l1.Val
		}

		if l2 != nil {
			nodesSum += l2.Val
		}

		res.Val = (nodesSum + draftNum) % 10
		draftNum = (nodesSum + draftNum) / 10

		preRes = res
		res.Next = &ListNode{Val: 0}
		res = res.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	var (
		tailList *ListNode
	)
	if l1 != nil {
		tailList = l1
	}

	if l2 != nil {
		tailList = l2
	}

	for tailList != nil {
		res.Val = (tailList.Val + draftNum) % 10
		draftNum = (tailList.Val + draftNum) / 10

		tailList = tailList.Next

		preRes = res
		res.Next = &ListNode{Val: 0}
		res = res.Next
	}

	switch draftNum {
	case 0:
		preRes.Next = nil
	default:
		res.Val = draftNum
	}

	return head
}

func addTwoNumbersRepeat(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		curSum, remainder int
	)
	res := l1

	for {
		curSum = l1.Val + l2.Val + remainder
		l1.Val, remainder = curSum%10, curSum/10

		if l1.Next == nil && l2.Next == nil {
			break
		}

		if l1.Next == nil && l2.Next != nil {
			l1.Next = &ListNode{}
		}

		if l2.Next == nil && l1.Next != nil {
			l2.Next = &ListNode{}
		}

		l1, l2 = l1.Next, l2.Next
	}

	if remainder != 0 {
		l1.Next = &ListNode{Val: remainder}
	}

	return res
}
