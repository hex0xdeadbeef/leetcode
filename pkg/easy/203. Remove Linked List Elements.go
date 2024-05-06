package easy

// https://leetcode.com/problems/remove-linked-list-elements/description/

import . "leetcode/pkg/datastructures/linkedlist"

func removeElements(head *ListNode, val int) *ListNode {

	var (
		prev, cur = &ListNode{Val: -1, Next: head}, head

		newHead = prev
	)

	for cur != nil {
		if cur.Val == val {
			for cur != nil && cur.Val == val {
				cur = cur.Next
			}
			prev.Next = cur
		}


		prev = prev.Next
		if prev == nil {
			break
		}
		cur = prev.Next

	}

	newHead = newHead.Next

	return newHead
}
