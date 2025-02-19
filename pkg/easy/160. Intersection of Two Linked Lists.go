package easy



import . "leetcode/pkg/datastructures/linkedlist"

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	const subtraction = 1e6

	var res *ListNode

	if headA == nil || headB == nil {
		return nil
	}

	curA, curB := headA, headB

	for curA != nil {
		curA.Val -= subtraction
		curA = curA.Next
	}

	for curB != nil {
		if curB.Val < 0 {
			res = curB
			break
		}

		curB = curB.Next
	}

	curA = headA
	for curA != nil {
		curA.Val += subtraction
		curA = curA.Next
	}

	return res
}

func getIntersectionNodeShameful(headA, headB *ListNode) *ListNode {
	var m = make(map[*ListNode]struct{}, 1<<8)

	if headA == nil || headB == nil {
		return nil
	}

	for headA != nil {
		m[headA], headA = struct{}{}, headA.Next
	}

	for headB != nil {
		if _, ok := m[headB]; ok {
			return headB
		}

		headB = headB.Next
	}

	return nil
}
