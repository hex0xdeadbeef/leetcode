package middle

// https://leetcode.com/problems/delete-leaves-with-a-given-value/?envType=daily-question&envId=2024-05-17

import . "leetcode/pkg/datastructures/treenode"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	var (
		preHead  = &TreeNode{Val: -1e9, Left: root, Right: nil}
		traverse func(root *TreeNode) bool
	)

	traverse = func(root *TreeNode) bool {
		if root.Left != nil && traverse(root.Left) {
			root.Left = nil
		}

		if root.Right != nil && traverse(root.Right) {
			root.Right = nil
		}

		return root.Left == nil && root.Right == nil && root.Val == target
	}

	traverse(preHead)

	if preHead.Left == nil {
		return nil
	}

	return preHead.Left
}
