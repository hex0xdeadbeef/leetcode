package easy

// https://leetcode.com/problems/linked-list-cycle/description/

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	const (
		padding = 1
		mark    = 10e5 + padding
	)

	for head != nil {

		if head.Val == mark {
			return true
		}

		head.Val = mark
		head = head.Next

	}

	return false
}
