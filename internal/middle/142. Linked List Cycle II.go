package middle

import . "leetcode/pkg/datastructures/linkedlist"

func detectCycle(head *ListNode) *ListNode {

	const sub = 1e6

	cur := head

	for cur != nil {
		if cur.Val < -1e5 {
			res := cur

			for cur.Val < -1e5 {
				cur.Val, cur = cur.Val+sub, cur.Next
			}

			return res
		}

		cur.Val, cur = cur.Val-sub, cur.Next
	}

	cur = head
	for cur != nil {
		cur.Val, cur = cur.Val+sub, cur.Next
	}

	return nil
}
