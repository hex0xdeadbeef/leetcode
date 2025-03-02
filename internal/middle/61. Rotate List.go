package middle

import (
	. "leetcode/pkg/datastructures/linkedlist"
)

func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	var n int
	newHead := head

	for newHead != nil {
		n++
		newHead = newHead.Next
	}
	newHead = head

	if k % n == 0 {
		return head
	}

	if n == 2 {
		if k%2 == 0 {
			return head
		}

		head.Next, head, head.Next.Next = nil, head.Next, head
		return head
	}

	for i := 0; newHead != nil; i++ {
		if (i+k)%n == 0 {
			break
		}
		newHead = newHead.Next
	}

	oldHead, cur := head, head
	for cur.Next != newHead {
		cur = cur.Next
	}

	cur.Next, cur = nil, newHead
	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = oldHead

	return newHead
}
