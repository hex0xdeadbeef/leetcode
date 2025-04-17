package middle

import . "leetcode/pkg/datastructures/treenode"

func sumNumbers(root *TreeNode) int {
	var res int

	var traverse func(root *TreeNode, curNum int)

	traverse = func(root *TreeNode, curNum int) {
		if root == nil {
			return
		}

		curNum += root.Val

		if root.Left == nil && root.Right == nil {
			res += curNum
		}

		traverse(root.Left, curNum*10)
		traverse(root.Right, curNum*10)
	}

	traverse(root, 0)

	return res
}
