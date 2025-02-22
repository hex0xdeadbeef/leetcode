package middle

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/

import (
	. "leetcode/pkg/datastructures/linkedlist"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head = reverseLinkedList(head)

	var (
		preNth = head
	)

	if head.Next == nil {
		return nil
	}

	if n == 1 {
		return reverseLinkedList(head.Next)
	}

	for n != 2 {
		preNth = preNth.Next
		n--
	}

	preNth.Next = preNth.Next.Next

	return reverseLinkedList(head)
}

func removeNthFromEndRepeat(head *ListNode, n int) *ListNode {
	head = reverseLinkedList(head)

	emptyNode := &ListNode{
		Next: head,
	}
	prev, cur := emptyNode, emptyNode

	for range n {
		prev, cur = cur, cur.Next
	}

	prev.Next = cur.Next

	return reverseLinkedList(emptyNode.Next)
}
