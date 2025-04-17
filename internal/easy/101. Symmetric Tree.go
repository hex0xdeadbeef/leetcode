package easy

import (
	. "leetcode/pkg/datastructures/treenode"
	"slices"
)

func isSymmetric(root *TreeNode) bool {
	var traverse func(*TreeNode, *[]int, bool)

	traverse = func(root *TreeNode, path *[]int, toLeft bool) {
		if root == nil {
			*path = append(*path, -1e6)
			return
		}

		*path = append(*path, root.Val)

		switch toLeft {
		case true:
			traverse(root.Left, path, toLeft)
		default:
			traverse(root.Right, path, toLeft)
		}

		switch toLeft {
		case true:
			traverse(root.Right, path, toLeft)
		default:
			traverse(root.Left, path, toLeft)
		}

	}

	s1, s2 := make([]int, 0), make([]int, 0)

	traverse(root.Left, &s1, true)
	traverse(root.Right, &s2, false)

	return slices.Equal(s1, s2)
}
