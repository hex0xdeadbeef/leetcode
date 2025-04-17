package easy

import . "leetcode/pkg/datastructures/treenode"

func preorderTraversal(root *TreeNode) []int {
	s := make([]int, 0)

	var traverse func(*TreeNode)

	traverse = func(root *TreeNode) {

		if root == nil {
			return
		}

		s = append(s, root.Val)

		traverse(root.Left)
		traverse(root.Right)
	}

	traverse(root)

	return s
}
