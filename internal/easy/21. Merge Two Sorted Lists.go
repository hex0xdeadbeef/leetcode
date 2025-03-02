package easy

// https://docs.google.com/spreadsheets/d/1DpWIdhlpni9LgS3oHYLu-V-Q2jxtcyQKBXgmcabauEE/edit?gid=0#gid=0

import . "leetcode/pkg/datastructures/linkedlist"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 != nil && list2 == nil {
		return list1
	}

	if list2 != nil && list1 == nil {
		return list2
	}

	emptyRoot := &ListNode{Val: -1}
	cur, cur1, cur2 := emptyRoot, list1, list2

	for cur1 != nil && cur2 != nil {
		switch {
		case cur1.Val <= cur2.Val:
			cur.Next, cur1 = cur1, cur1.Next
		default:
			cur.Next, cur2 = cur2, cur2.Next
		}

		cur = cur.Next
	}

	if cur2 != nil {
		cur1 = cur2
	}

	for cur1 != nil {
		cur.Next, cur1 = cur1, cur1.Next
		cur = cur.Next
	}

	return emptyRoot.Next

}
