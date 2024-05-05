package middle

// https://leetcode.com/problems/delete-node-in-a-linked-list/description/?envType=daily-question&envId=2024-05-05

import . "leetcode/pkg/datastructures/linkedlist"

func deleteNode(node *ListNode) {

	var (
		prevOfValToDelete *ListNode

		prev, cur *ListNode = node, node.Next
	)

	for cur != nil {
		prevOfValToDelete = prev
		prev.Val, prev, cur = cur.Val, cur, cur.Next
	}

	prevOfValToDelete.Next = nil

}
