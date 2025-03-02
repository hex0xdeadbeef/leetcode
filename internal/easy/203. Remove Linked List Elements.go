package easy

// https://leetcode.com/problems/remove-linked-list-elements/description/

import . "leetcode/pkg/datastructures/linkedlist"

func removeElements(head *ListNode, val int) *ListNode {

	var (
		prev, cur = &ListNode{Val: -1, Next: head}, head

		newHead = prev
	)

	for cur != nil {
		if cur.Val == val {
			for cur != nil && cur.Val == val {
				cur = cur.Next
			}
			prev.Next = cur
		}

		prev = prev.Next
		if prev == nil {
			break
		}
		cur = prev.Next

	}

	newHead = newHead.Next

	return newHead
}

func removeElementsCheatedA(head *ListNode, val int) *ListNode {
	var (
		s   = make([]*ListNode, 0, 1<<14)
		cur = head
	)

	for cur != nil {
		s = append(s, cur)
		next := cur.Next

		cur.Next = nil
		cur = next
	}

	for i, v := range s {
		if v.Val == val {
			s[i] = nil
		}
	}

	var (
		res = &ListNode{Val: -1, Next: nil}
	)
	cur = res

	for i := 0; i < len(s); i++ {
		if s[i] != nil {
			cur.Next = s[i]
			cur = cur.Next
		}
	}

	if res.Next == nil {
		return nil
	}

	return res.Next
}

func removeElementsCheatedB(head *ListNode, val int) *ListNode {
	var (
		res     = &ListNode{Val: -1e6, Next: nil}
		resTail = res

		cur = head
	)

	for cur != nil {
		switch {
		case cur.Val == val:

		default:
			resTail.Next = cur
			resTail = resTail.Next
		}

		cur = cur.Next
	}

	if resTail.Next != nil && resTail.Next.Val == val {
		resTail.Next = nil
	}

	if res.Next == nil {
		return nil
	}

	return res.Next
}

func removeElementsRepeat(head *ListNode, val int) *ListNode {
	emptyRoot := &ListNode{
		Val: -1,
	}
	root := emptyRoot

	for {
		for head != nil && head.Val == val {
			head = head.Next
		}

		root.Next = head
		root = root.Next

		if head == nil {
			break
		}

		head = head.Next
	}

	return emptyRoot.Next
}


