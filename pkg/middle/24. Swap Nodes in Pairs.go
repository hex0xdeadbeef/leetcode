package middle

import (
	. "leetcode/pkg/datastructures/linkedlist"
)

func SwapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	l, r := head, head.Next
	nodeToBindTo := &ListNode{}
	root := nodeToBindTo

	for {
		if l.Next == nil {
			break
		}

		l.Next, r.Next = r.Next, l
		nodeToBindTo.Next, nodeToBindTo = r, l

		if l.Next == nil {
			break
		}

		l, r = l.Next, l.Next.Next

	}

	return root.Next
}
