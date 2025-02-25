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

func oddEvenListRepeat(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var oddRoot, evenRoot, cur *ListNode
	odd, even := &ListNode{}, &ListNode{}

	for i := 1; head != nil; i++ {
		cur = head

		switch i % 2 {
		case 1:
			if i == 1 {
				oddRoot = cur
			}

			odd.Next, odd = cur, cur
		default:
			if i == 2 {
				evenRoot = cur
			}

			even.Next, even = cur, cur
		}

		head.Next, head = nil, head.Next
	}
	odd.Next = evenRoot

	return oddRoot
}
