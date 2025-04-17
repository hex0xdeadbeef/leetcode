package easy

import . "leetcode/pkg/datastructures/treenode"

func isBalanced(root *TreeNode) bool {
	var traverse func(*TreeNode)

	isBalanced := true

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		dL := findDepth(root.Left, 0)
		dR := findDepth(root.Right, 0)

		if dif := abs(dL - dR); dif > 1 && isBalanced {
			isBalanced = !isBalanced
		}

		traverse(root.Left)
		traverse(root.Right)
	}

	traverse(root)

	return isBalanced
}

func findDepth(root *TreeNode, d int) int {
	if root == nil {
		return d
	}

	d++

	return max(findDepth(root.Left, d), findDepth(root.Right, d))
}
