package easy

// https://leetcode.com/problems/reverse-linked-list/?envType=study-plan-v2&envId=leetcode-75

import (
	. "leetcode/pkg/datastructures/linkedlist"
)

func reverseList(head *ListNode) *ListNode {

	var (
		a, b, c *ListNode
	)

	switch {
	case head == nil:
		return nil

	case head.Next == nil:
		return head
	}

	a, b, c = nil, head, head

	for c.Next != nil {
		c = c.Next
		b.Next = a
		a = b
		b = c
	}

	b.Next = a

	return b

}

func reverseListSimple(head *ListNode) *ListNode {
	var (
		a, b, c *ListNode = nil, head, head.Next
	)

	if b == nil {
		return head
	}

	for c != nil {
		b.Next, a, b, c = a, b, c, c.Next
	}

	b.Next = a

	return b
}
