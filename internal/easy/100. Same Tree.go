package easy

import (
	. "leetcode/pkg/datastructures/treenode"
	"slices"
)

func isSameTree(t1 *TreeNode, t2 *TreeNode) bool {
	var traverse func(*TreeNode, *[]int)

	traverse = func(root *TreeNode, path *[]int) {
		if root == nil {
			*path = append(*path, -1e6)
			return
		}

		*path = append(*path, root.Val)

		traverse(root.Left, path)
		traverse(root.Right, path)
	}

	s1, s2 := make([]int, 0), make([]int, 0)

	traverse(t1, &s1)
	traverse(t2, &s2)

	return slices.Equal(s1, s2)
}
