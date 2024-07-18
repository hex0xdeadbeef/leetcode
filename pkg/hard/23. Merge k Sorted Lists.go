package hard

// https://leetcode.com/problems/merge-k-sorted-lists/description/

import (
	. "leetcode/pkg/datastructures/linkedlist"
	"sort"
)

type Nodes []*ListNode

func (n *Nodes) Len() int           { return len(*n) }
func (n *Nodes) Swap(i, j int)      { (*n)[i], (*n)[j] = (*n)[j], (*n)[i] }
func (n *Nodes) Less(i, j int) bool { return (*n)[i].Val < (*n)[j].Val }

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
	sort.Sort(&nodes)

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
