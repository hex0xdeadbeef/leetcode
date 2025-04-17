package easy

// https://leetcode.com/problems/leaf-similar-trees/submissions/1226550945/?envType=study-plan-v2&envId=leetcode-75

import (
	. "leetcode/pkg/datastructures/treenode"
	"slices"
)

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var s1, s2 []int

	var traverse func(*TreeNode, *[]int)

	traverse = func(root *TreeNode, s *[]int) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			*s = append(*s, root.Val)
			return
		}

		traverse(root.Left, s)
		traverse(root.Right, s)
	}

	traverse(root1, &s1)
	traverse(root2, &s2)

	return slices.Equal(s1, s2)
}
