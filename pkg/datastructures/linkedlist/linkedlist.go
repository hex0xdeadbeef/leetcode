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
