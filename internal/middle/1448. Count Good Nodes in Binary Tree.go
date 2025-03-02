package middle

// https://leetcode.com/problems/count-good-nodes-in-binary-tree/?envType=study-plan-v2&envId=leetcode-75

import (
	. "leetcode/pkg/datastructures/treenode"
)

func goodNodes(root *TreeNode) int {
	var (
		count = 0

		traverse func(root *TreeNode, curMaxNum int)
	)

	traverse = func(root *TreeNode, curMaxNum int) {

		if root == nil {
			return
		}

		if curMaxNum <= root.Val {
			curMaxNum = root.Val
			count++
		}

		traverse(root.Left, curMaxNum)
		traverse(root.Right, curMaxNum)

	}

	traverse(root, -10e5)

	return count
}
