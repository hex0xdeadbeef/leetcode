package middle

// https://leetcode.com/problems/path-sum-iii/?envType=study-plan-v2&envId=leetcode-75

import (
	. "leetcode/pkg/datastructures/treenode"
)

func pathSum(root *TreeNode, targetSum int) int {
	const (
		zeroSum = 0
	)

	var (
		count int

		traverse      func(root *TreeNode)
		getSubtreeSum func(root *TreeNode, curSum int)
	)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		getSubtreeSum(root, zeroSum)

		traverse(root.Left)
		traverse(root.Right)

	}

	getSubtreeSum = func(root *TreeNode, curSum int) {
		if root == nil {
			return
		}

		newSum := root.Val + curSum
		if newSum == targetSum {
			count++
		}

		getSubtreeSum(root.Left, newSum)
		getSubtreeSum(root.Right, newSum)

	}

	traverse(root)

	return count
}
