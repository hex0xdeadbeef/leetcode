package easy

// https://leetcode.com/problems/linked-list-cycle/description/

import . "leetcode/pkg/datastructures/linkedlist"

func hasCycleA(head *ListNode) bool {
	const (
		mark = 1e6
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

func hasCycleB(head *ListNode) bool {
	var (
		slow = head
		fast = head
	)

	for slow != nil && fast != nil {
		slow = slow.Next

		if fast.Next != nil {
			fast = fast.Next.Next
		}

		if slow == fast && (slow.Next != nil && fast.Next != nil) {
			return true
		}
	}

	return false

}

func hasCycleMapRepeat(head *ListNode) bool {
	var m = make(map[*ListNode]struct{}, 1<<8)

	if head == nil {
		return false
	}

	for {
		_, ok := m[head]
		if ok {
			return true
		}

		m[head] = struct{}{}

		head = head.Next
		if head == nil {
			return false
		}
	}
}

func hasCycleFastSlowRepeat(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head

	for {
		if fast == nil || fast.Next == nil || slow == nil {
			return false
		}

		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		}

		if fast == slow {
			return true
		}
	}

}
