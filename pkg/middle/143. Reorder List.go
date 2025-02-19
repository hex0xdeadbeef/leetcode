package middle

// https://leetcode.com/problems/reorder-list/description/

import (
	. "leetcode/pkg/datastructures/linkedlist"
)

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

func reorderListRepeat(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	l, m, r := SplitLinkedList(head)
	r = ReverseLinkedList(r)

	empty := &ListNode{}
	cur := empty

	for l != nil {
		cur.Next, l = l, l.Next
		cur = cur.Next

		cur.Next, r = r, r.Next
		if cur.Next != nil {
			cur = cur.Next
		}
	}

	cur.Next = m
}
