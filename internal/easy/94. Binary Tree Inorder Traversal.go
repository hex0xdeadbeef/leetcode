package easy

import . "leetcode/pkg/datastructures/treenode"

func inorderTraversal(root *TreeNode) []int {
	s := make([]int, 0)

	var traverse func(*TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		traverse(root.Left)

		s = append(s, root.Val)

		traverse(root.Right)
	}

	traverse(root)

	return s
}
