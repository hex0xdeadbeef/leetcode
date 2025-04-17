package middle

import . "leetcode/pkg/datastructures/treenode"

func flatten(root *TreeNode) {
	var s []*TreeNode

	var traverse func(*TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		s = append(s, root)
		traverse(root.Left)
		traverse(root.Right)
	}

	traverse(root)

	for i := range s {
		s[i].Left, s[i].Right = nil, nil
	}

	for i := range len(s) - 1 {
		s[i].Right = s[i+1]
	}


}
