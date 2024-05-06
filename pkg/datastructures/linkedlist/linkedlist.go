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

/* node1 := &ListNode{Val: 1, Next: nil}
node2 := &ListNode{Val: 1, Next: nil}
node3 := &ListNode{Val: 1, Next: nil}
node4 := &ListNode{Val: 1, Next: nil}
node5 := &ListNode{Val: 1, Next: nil}
node6 := &ListNode{Val: 1, Next: nil}
node7 := &ListNode{Val: 1, Next: nil}
node8 := &ListNode{Val: 1, Next: nil}

node1.Next,
	node2.Next,
	node3.Next,
	node4.Next,
	node5.Next,
	node6.Next,
	node7.Next =

	node2,
	node3,
	node4,
	node5,
	node6,
	node7,
	node8

res := Func(node1, 6)

for res != nil {
	fmt.Printf("%d ", res.Val)
	res = res.Next
} */
