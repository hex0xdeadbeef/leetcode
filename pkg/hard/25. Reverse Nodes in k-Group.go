package hard

// https://leetcode.com/problems/reverse-nodes-in-k-group/

import (
	. "leetcode/pkg/datastructures/linkedlist"
)

func ReverseKGroup(head *ListNode, k int) *ListNode {
	var (
		nodes         = DelinkList(head)
		wholePartsCnt = len(nodes) / k
	)

	for i := 0; i < len(nodes); i += k {
		if wholePartsCnt == 0 {
			break
		}

		ReverseElems(nodes[i : i+k])
		wholePartsCnt--
	}

	for i, node := range nodes {
		if i+1 == len(nodes) {
			break
		}

		node.Next = nodes[i+1]
	}

	return nodes[0]
}

func ReverseElems[T any](s []T) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func DelinkList(list *ListNode) []*ListNode {
	var (
		nodes = make([]*ListNode, 0, 1<<8)
	)

	for list != nil {
		nodes = append(nodes, list)

		list, list.Next = list.Next, nil
	}

	return nodes
}

func ReverseKGroupRepeat(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	beforeL := &ListNode{Next: head}
	root := beforeL

	var hopsCount int
	head = beforeL
	for {
		for hopsCount = 0; hopsCount < k && head.Next != nil; head, hopsCount = head.Next, hopsCount+1 {
		}

		if hopsCount < k {
			break
		}

		l, rAfter := beforeL.Next, head.Next
		head.Next, head, beforeL.Next = nil, head.Next, nil

		l = ReverseLinkedList(l)
		beforeL.Next = l

		for ; l.Next != nil; l = l.Next {
		}
		l.Next = rAfter

		beforeL, head = l, l
	}

	return root.Next

}
