package easy

import (
	. "leetcode/pkg/datastructures/treenode"
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
	var s []string

	var traverse func(*TreeNode, string)

	traverse = func(root *TreeNode, path string) {
		if root == nil {
			return
		}

		path += strconv.Itoa(root.Val)

		if root.Left == nil && root.Right == nil {
			s = append(s, path)
		}

		if root.Left != nil {
			traverse(root.Left, path+"->")
		}

		if root.Right != nil {
			traverse(root.Right, path+"->")
		}
	}

	traverse(root, "")

	return s
}
