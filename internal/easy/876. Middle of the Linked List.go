package easy

// https://leetcode.com/problems/middle-of-the-linked-list/description/?envType=daily-question&envId=2024-06-02

import (
	. "leetcode/pkg/datastructures/linkedlist"
)

func middleNode(head *ListNode) *ListNode {
	var (
		cnt int

		slow, fast *ListNode = head, head

		firstMiddle *ListNode
	)

	for fast.Next != nil && fast.Next.Next != nil {
		cnt++

		slow = slow.Next
		fast = fast.Next.Next
	}

	firstMiddle = slow

	for slow.Next != nil {
		cnt++
		slow = slow.Next
	}

	if cnt%2 == 1 {
		return firstMiddle.Next
	}
	return firstMiddle.Next
}

func middleNodeRepeat(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	emptyNode := &ListNode{
		Next: head,
	}

	slow, fast := emptyNode, emptyNode

	for {
		if fast == nil {
			return slow
		}

		if fast.Next == nil {
			return slow.Next
		}

		slow, fast = slow.Next, fast.Next.Next
	}
}
