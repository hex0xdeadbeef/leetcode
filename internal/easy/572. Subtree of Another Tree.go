package easy

import . "leetcode/pkg/datastructures/treenode"

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var res bool

	var isSameTree func(root1, root2 *TreeNode) bool

	isSameTree = func(root1, root2 *TreeNode) bool {
		if root1 == nil && root2 == nil {
			return true
		}

		if root1 == nil || root2 == nil {
			return false
		}

		if root1.Val != root2.Val {
			return false
		}

		return isSameTree(root1.Left, root2.Left) && isSameTree(root1.Right, root2.Right)
	}

	var traverse func(*TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		if res {
			return
		}

		if isSameTree(root, subRoot) {
			res = true
			return
		}

		traverse(root.Left)
		traverse(root.Right)
	}

	traverse(root)

	return res
}
