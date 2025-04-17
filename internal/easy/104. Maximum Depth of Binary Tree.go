package easy

import . "leetcode/pkg/datastructures/treenode"

func maxDepth(root *TreeNode) int {
	var resDepth int

	var traverse func(*TreeNode, int)

	traverse = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		if depth > resDepth {
			resDepth = depth
		}

		traverse(root.Left, depth+1)
		traverse(root.Right, depth+1)
	}

	traverse(root, 1)

	return resDepth
}
