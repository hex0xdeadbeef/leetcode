package middle

import (
	. "leetcode/pkg/datastructures/linkedlist"
)

func doubleIt(head *ListNode) *ListNode {

	const (
		doubler = 2
		ten     = 10
	)

	var (
		left int

		rightHead *ListNode = reverseLinkedList(head)

		cur *ListNode = rightHead

		doubledNodeVal int
	)

	for cur != nil {
		doubledNodeVal = getDoubledNodeValue(cur, left)

		cur.Val, left = doubledNodeVal%ten, doubledNodeVal/ten

		if cur.Next == nil && left != 0 {
			cur.Next = &ListNode{Val: left, Next: nil}
			cur = cur.Next
			break
		}
		cur = cur.Next
	}

	return reverseLinkedList(rightHead)

}

func getDoubledNodeValue(node *ListNode, left int) int {
	const (
		doubler = 2
	)

	return node.Val*doubler + left
}

func reverseLinkedList(head *ListNode) *ListNode {
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
