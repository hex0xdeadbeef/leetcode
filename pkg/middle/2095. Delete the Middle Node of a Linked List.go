package middle

// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/description/?envType=study-plan-v2&envId=leetcode-75

import . "leetcode/pkg/datastructures/linkedlist"

func deleteMiddle(head *ListNode) *ListNode {

	var (
		slow, fast = head, head

		prev *ListNode
	)

	if head.Next == nil {
		return nil
	}

	for {

		prev = slow

		slow, fast = slow.Next, fast.Next.Next

		if fast == nil || fast.Next == nil {
			break
		}
	}

	prev.Next = slow.Next

	return head
}
