package middle

// https://leetcode.com/problems/all-elements-in-two-binary-search-trees/description/

import (
	. "leetcode/pkg/datastructures/treenode"
)

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var (
		s1, s2            = make([]int, 0, 1<<6), make([]int, 0, 1<<6)
		traverseAndAppend func(node *TreeNode, s []int) []int
	)
	traverseAndAppend = func(node *TreeNode, s []int) []int {
		if node == nil {
			return s
		}

		if node.Left == nil && node.Right == nil {
			return append(s, node.Val)
		}

		if node.Left != nil {
			s = traverseAndAppend(node.Left, s)
		}

		s = append(s, node.Val)

		if node.Right != nil {
			s = traverseAndAppend(node.Right, s)
		}

		return s
	}

	return mergeSortedSlices(traverseAndAppend(root1, s1), traverseAndAppend(root2, s2))
}

func mergeSortedSlices(s1, s2 []int) []int {
	var (
		res = make([]int, 0, len(s1)+len(s2))

		p1, p2 int
	)

	for p1 < len(s1) && p2 < len(s2) {
		v1, v2 := s1[p1], s2[p2]

		switch v1 < v2 {
		case true:
			res = append(res, v1)
			p1++
		default:
			res = append(res, v2)
			p2++
		}
	}

	if p1 == len(s1) && p2 == len(s2) {
		return res
	}

	var (
		tail []int
		p    int
	)

	switch p1 < len(s1) {
	case true:
		tail = s1
		p = p1

	default:
		tail = s2
		p = p2
	}

	for ; p < len(tail); p++ {
		res = append(res, tail[p])
	}

	return res
}
