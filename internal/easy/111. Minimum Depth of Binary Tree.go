package easy

import . "leetcode/pkg/datastructures/treenode"

func minDepth(root *TreeNode) int {
	minD := 1_000

	var traverse func(*TreeNode, int)

	traverse = func(root *TreeNode, d int) {
		if root == nil {
			if d < minD {
				minD = d
			}
			return
		}

		d++

		traverse(root.Left, d)
		traverse(root.Right, d)
	}

	traverse(root, -1)

	return minD
}
