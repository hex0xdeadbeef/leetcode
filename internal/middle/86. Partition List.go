package middle

import . "leetcode/pkg/datastructures/linkedlist"

func listPartition(head *ListNode, x int) *ListNode {
	l, eg := []*ListNode(nil), []*ListNode(nil)

	cur := head

	for ; cur != nil; cur = cur.Next {
		if cur.Val < x {
			l = append(l, cur)
			continue
		}

		eg = append(eg, cur)
	}

	l = referenceNodes(append(l, eg...))

	if len(l) != 0 {
		return l[0]
	}

	return nil
}

func referenceNodes(s []*ListNode) []*ListNode {
	for i := range len(s) - 1 {
		s[i].Next = s[i+1]
	}

	if len(s) != 0 {
		s[len(s)-1].Next = nil
	}

	return s
}
