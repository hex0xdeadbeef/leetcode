package middle

// https://leetcode.com/problems/reorder-list/description/

import . "leetcode/pkg/datastructures/linkedlist"

func reorderList(head *ListNode) {
	var (
		savedHead = head
		list      = getCopiedList(head.Next)
		l, r      = 0, len(list) - 1
	)

	for i := 0; i < len(list); i++ {
		switch i % 2 {
		case 0:
			head.Next = list[r]
			r--
		case 1:
			head.Next = list[l]
			l++
		}

		head = head.Next
	}
	head.Next = nil

	head = savedHead
}

func getCopiedList(root *ListNode) []*ListNode {
	var (
		list = make([]*ListNode, 0, 1<<10)
	)

	for root != nil {
		list = append(list, root)
		root = root.Next
	}

	return list
}
