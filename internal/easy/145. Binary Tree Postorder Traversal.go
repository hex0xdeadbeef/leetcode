package easy

import . "leetcode/pkg/datastructures/treenode"

func postorderTraversal(root *TreeNode) []int {
	s := make([]int, 0)

	var traverse func(*TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		traverse(root.Left)
		traverse(root.Right)

		s = append(s, root.Val)

	}

	traverse(root)


	return s
}
