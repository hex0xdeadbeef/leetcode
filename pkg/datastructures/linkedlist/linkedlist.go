package linkedlist

import "fmt"

type ListNode struct {
	Next *ListNode
	Val  int
}

func NewLinkedList(nums []int) *ListNode {
	head := &ListNode{Next: nil, Val: nums[0]}

	var prev *ListNode = head

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
