package easy

// https://leetcode.com/problems/evaluate-boolean-binary-tree/description/?envType=daily-question&envId=2024-05-16

import . "leetcode/pkg/datastructures/treenode"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func evaluateTree(root *TreeNode) bool {
	const (
		FALSE, TRUE = 0, 1
		OR, AND     = 2, 3
	)

	var (
		traverse func(root *TreeNode) bool
	)
	traverse = func(root *TreeNode) bool {
		if root.Left == nil && root.Right == nil {
			if root.Val == TRUE {
				return true
			}

			return false
		}

		if root.Val == OR {
			return traverse(root.Left) || traverse(root.Right)
		}

		if root.Val == AND {
			return traverse(root.Left) && traverse(root.Right)
		}

		panic("UNREACHABLE")
	}

	return traverse(root)

}
