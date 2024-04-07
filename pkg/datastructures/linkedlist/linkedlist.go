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
