package middle

import . "leetcode/pkg/datastructures/treenode"

func isValidBST(root *TreeNode) bool {
	res := true

	var validate func(root, subRoot *TreeNode) bool

	validate = func(root, subRoot *TreeNode) bool {
		if subRoot == nil {
			return true
		}

		l, r := true, true

		if subRoot.Left != nil {
			l = subRoot.Left.Val < subRoot.Val && subRoot.Left.Val < root.Val
		}

		if subRoot.Right != nil {
			r = subRoot.Right.Val > subRoot.Val && subRoot.Right.Val > root.Val
		}

		lSub := validate(root, subRoot.Left)
		rSub := validate(root, subRoot.Right)

		return l && r && lSub && rSub
	}

	var traverse func(*TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Left != nil {
			if !validate(root, root.Left) {
				res = false
			}
		}

		if root.Right != nil {
			if !validate(root, root.Right) {
				res = false
			}
		}

		traverse(root.Left)
		traverse(root.Right)

	}

	traverse(root)

	return res
}
