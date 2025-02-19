package easy

import (
	. "leetcode/pkg/datastructures/linkedlist"
)

func IsListPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	if head.Next == nil {
		return true
	}

	l, _, r := SplitLinkedList(head)
	r = reverseList(r)

	if l == nil || r == nil {
		return false
	}

	for {
		if r == nil && l == nil {
			return true
		}

		if l == nil || r == nil {
			return false
		}

		if r.Val != l.Val {
			return false
		}

		l, r = l.Next, r.Next
	}
}
