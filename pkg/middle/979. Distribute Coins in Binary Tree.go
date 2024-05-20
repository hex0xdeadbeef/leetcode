package middle

// https://leetcode.com/problems/distribute-coins-in-binary-tree/description/

import . "leetcode/pkg/datastructures/treenode"

func distributeCoins(root *TreeNode) int {
	var (
		traverse func(parent, root *TreeNode)

		totalSteps int
	)

	traverse = func(parent, root *TreeNode) {
		if root == nil {
			return
		}

		traverse(root, root.Left)
		traverse(root, root.Right)

		totalSteps += abs(root.Val - 1)

		if parent != nil {
			parent.Val += root.Val - 1
		}
	}

	traverse(nil, root)

	return totalSteps

}

// func abs(n int) int {
// 	if n < 0 {
// 		return -1 * n
// 	}

// 	return n
// }
