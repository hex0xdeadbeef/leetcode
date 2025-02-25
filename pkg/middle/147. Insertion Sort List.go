package middle

import (
	. "leetcode/pkg/datastructures/linkedlist"
)

func InsertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == nil {
		return head
	}

	empty, cur := &ListNode{Val: -1e6}, head
	sortedCurPrev, sortedCur := empty, empty

	for cur != nil {
		switch {
		case sortedCur.Next != nil:
			for sortedCur != nil && cur.Val > sortedCur.Val {
				sortedCurPrev, sortedCur = sortedCur, sortedCur.Next
			}

			sortedCurPrev.Next, cur.Next, cur = cur, sortedCur, cur.Next

		default:
			sortedCur.Next = cur
			cur, cur.Next = cur.Next, nil
		}

		sortedCurPrev, sortedCur = empty, empty
	}

	return empty.Next
}
