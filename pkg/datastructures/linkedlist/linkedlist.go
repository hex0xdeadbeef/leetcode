package linkedlist

import "fmt"

type ListNode struct {
	Next *ListNode
	Val  int
}

func NewLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Next: nil, Val: nums[0]}

	prev := head

	for i := 1; i < len(nums); i++ {
		node := &ListNode{Next: nil, Val: nums[i]}
		prev.Next = node
		prev = node
	}

	return head
}

func (l *ListNode) String() string {
	var (
		result = ""
	)

	for l != nil {
		result += fmt.Sprintf("%d -> ", l.Val)
		l = l.Next
	}

	result += "nil"

	return result
}

func NewList() *ListNode {
	n1, n2, n3, n4, n5 := &ListNode{Val: 1}, &ListNode{Val: 2}, &ListNode{Val: 3}, &ListNode{Val: 4}, &ListNode{Val: 5}

	n1.Next, n2.Next, n3.Next, n4.Next = n2, n3, n4, n5

	return n1
}

func SplitLinkedList(head *ListNode) (left, middle, right *ListNode) {
	if head == nil {
		return nil, nil, nil
	}

	emptyRoot := &ListNode{
		Next: head,
	}

	slow, fast := emptyRoot, emptyRoot
	var prevSlow *ListNode

	for {
		prevSlow, slow = slow, slow.Next
		if fast != nil && fast.Next != nil {
			fast = fast.Next.Next
		}

		if fast == nil {
			right, middle, slow.Next, left, prevSlow.Next = slow.Next, slow, nil, emptyRoot.Next, nil
			return
		}

		if fast.Next == nil {
			right, left, slow.Next = slow.Next, emptyRoot.Next, nil
			return
		}
	}
}

func ReverseLinkedList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var (
		a, b, c *ListNode = nil, head, head.Next
	)

	if b == nil {
		return head
	}

	for c != nil {
		b.Next, a, b, c = a, b, c, c.Next
	}

	b.Next = a

	return b
}

func MergeSortedLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 != nil && list2 == nil {
		return list1
	}

	if list2 != nil && list1 == nil {
		return list2
	}

	emptyRoot := &ListNode{Val: -1}
	cur, cur1, cur2 := emptyRoot, list1, list2

	for cur1 != nil && cur2 != nil {
		switch {
		case cur1.Val <= cur2.Val:
			cur.Next, cur1 = cur1, cur1.Next
		default:
			cur.Next, cur2 = cur2, cur2.Next
		}

		cur = cur.Next
	}

	if cur2 != nil {
		cur1 = cur2
	}

	for cur1 != nil {
		cur.Next, cur1 = cur1, cur1.Next
		cur = cur.Next
	}

	return emptyRoot.Next

}

func SortListMerge(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	l, m, r := SplitLinkedList(head)
	if m != nil {
		m.Next, r = r, m
	}

	return MergeSortedLists(SortListMerge(l), SortListMerge(r))
}
