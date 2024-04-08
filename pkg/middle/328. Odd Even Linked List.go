package middle

// https://leetcode.com/problems/odd-even-linked-list/?envType=study-plan-v2&envId=leetcode-75

import (
	. "leetcode/pkg/datastructures/linkedlist"
)

func oddEvenList(head *ListNode) *ListNode {

	var (
		odd, even *ListNode

		evensHead, evensCur *ListNode
	)

	if head == nil || head.Next == nil {
		return head
	}

	odd, even = head, head.Next

	// Pre-initialization
	evensHead, evensCur = even, even

	even = even.Next

	odd.Next = even

	evensCur.Next = nil

	for {
		if even == nil || odd == nil {
			break
		}
		odd = odd.Next
		even = even.Next

		evensCur.Next = even
		evensCur = evensCur.Next
		if even == nil {
			break
		}
		even = even.Next
		odd.Next = even
		evensCur.Next = nil

		if even == nil {
			break
		}

	}

	if odd == nil {
		head.Next = evensHead
		return head
	}
	odd.Next = evensHead

	return head
}
