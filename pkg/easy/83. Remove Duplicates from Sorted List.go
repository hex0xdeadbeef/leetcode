package easy

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/

import . "leetcode/pkg/datastructures/linkedlist"

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	var (
		res = head

		cur = head
	)

	for {
		if cur == nil {
			return res
		}

		for cur.Next != nil && cur.Val == cur.Next.Val {
			if cur.Next == nil {
				return res
			}

			cur.Next = cur.Next.Next

		}

		cur = cur.Next

	}
}
