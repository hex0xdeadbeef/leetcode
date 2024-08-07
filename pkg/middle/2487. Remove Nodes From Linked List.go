package middle

// https://leetcode.com/problems/remove-nodes-from-linked-list/description/?envType=daily-question&envId=2024-05-06

import (
	. "leetcode/pkg/datastructures/linkedlist"
)

func removeNodes(head *ListNode) *ListNode {

	head = reverseLinkedList(head)

	var (
		base, runner *ListNode = head, head.Next
	)

outer:
	for {
		for runner != nil && runner.Val < base.Val {
			runner = runner.Next
		}

		if runner == nil {
			base.Next = nil
			break outer
		}

		base.Next = runner
		base = runner
		runner = base.Next
	}

	head = reverseLinkedList(head)

	return head
}
