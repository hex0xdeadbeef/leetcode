package middle

import . "leetcode/pkg/datastructures/linkedlist"

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil || left == right {
		return head
	}

	var lPrev, l, rAfter *ListNode
	empty := &ListNode{Next: head}
	prev, cur := empty, empty

Out:
	for i := 0; cur != nil; i++ {
		switch i {
		case left:
			prev.Next, lPrev, l = nil, prev, cur

		case right:
			cur.Next, rAfter = nil, cur.Next
			break Out
		}

		prev, cur = cur, cur.Next
	}

	lPrev.Next = reverseLinkedList(l)

	cur = lPrev
	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = rAfter

	return empty.Next
}
