package middle

import (
	. "leetcode/pkg/datastructures/linkedlist"
)

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	pivot := head
	head.Next, head = nil, head.Next

	l, r := dummyPartition(pivot.Val, head)
	l, r = sortList(l), sortList(r)
	lRoot := l

	for l != nil && l.Next != nil {
		l = l.Next
	}

	pivot.Next = r
	if lRoot != nil {
		l.Next, pivot.Next = pivot, r
		return lRoot
	}

	return pivot
}

func dummyPartition(pivotValue int, node *ListNode) (lPart, rPart *ListNode) {
	if node == nil {
		return nil, nil
	}

	if node.Next == nil {
		if node.Val <= pivotValue {
			return node, nil
		}

		return nil, node
	}

	lPart, rRoot := new(ListNode), new(ListNode)
	l, r := lPart, rRoot

	for ; node != nil; node.Next, node = nil, node.Next {
		if node.Val <= pivotValue {
			l.Next, l = node, node
			continue
		}
		r.Next, r = node, node

	}

	return lPart.Next, rRoot.Next
}

func sortListMerge(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	l, m, r := SplitLinkedList(head)
	if m != nil {
		m.Next, r = r, m
	}

	return MergeSortedLists(sortListMerge(l), sortListMerge(r))
}
