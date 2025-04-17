package easy

import . "leetcode/pkg/datastructures/treenode"

func hasPathSum(root *TreeNode, targetSum int) bool {
	var traverse func(*TreeNode, int)

	var isSumMet bool

	traverse = func(root *TreeNode, curSum int) {
		if root == nil {
			return
		}

		curSum += root.Val

		if root.Left == nil && root.Right == nil && curSum == targetSum && !isSumMet {
			isSumMet = !isSumMet
		}

		traverse(root.Left, curSum)
		traverse(root.Right, curSum)
	}

	if root != nil {
		traverse(root, 0)
	}

	return isSumMet
}
