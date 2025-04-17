package middle

import . "leetcode/pkg/datastructures/treenode"

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return p
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	if l != nil && r != nil {
		return root
	}

	if l != nil {
		return l
	}

	return r
}
