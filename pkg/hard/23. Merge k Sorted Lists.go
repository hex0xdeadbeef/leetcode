package hard

// https://leetcode.com/problems/merge-k-sorted-lists/description/

import (
	. "leetcode/pkg/datastructures/linkedlist"
	"slices"
)

type Nodes []*ListNode

func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	var (
		nodesToSort = make([]*ListNode, 0, 1<<11)
	)

	for i := range lists {
		if sigleListNodes := prepareNodes(lists[i]); sigleListNodes != nil {
			nodesToSort = append(nodesToSort, sigleListNodes...)
		}
	}
	nodes := Nodes(nodesToSort)
	slices.SortFunc(nodesToSort, func(a, b *ListNode) int {
		return a.Val - b.Val
	})

	for i, n := range nodes {
		if i+1 < len(nodes) {
			n.Next = nodes[i+1]
		}
	}

	if len(nodes) == 0 {
		return nil
	}

	return nodes[0]
}

func prepareNodes(list *ListNode) []*ListNode {
	if list == nil {
		return nil
	}

	var (
		nodes = make([]*ListNode, 0, 1<<8)
		cur   = list
	)

	for cur != nil {
		nodes = append(nodes, cur)
		cur, cur.Next = cur.Next, nil
	}

	return nodes
}

func mergeKLists(lists []*ListNode) *ListNode {
	empty := &ListNode{}
	cur := empty

	for _, v := range lists {
		if v == nil {
			continue
		}

		cur.Next = v
		for ; cur.Next != nil; cur = cur.Next {
		}
	}

	return SortListMerge(empty.Next)
}
