package middle

import (
	. "leetcode/pkg/datastructures/linkedlist"
)

func DeleteDuplicates(head *ListNode) *ListNode {
	const bigNegativeNum = -1e6
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	headPrev := &ListNode{Next: head}

	lPrev, l, r := headPrev, head, head

	var d int
	for r != nil {
		d = 0

		for r != nil && l.Val == r.Val {
			r = r.Next
			d++
		}

		if d == 1 {
			lPrev, l = l, r
			continue
		}

		lPrev.Next, l = r, r
	}

	return headPrev.Next
}
