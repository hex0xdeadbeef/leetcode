package easy

// https://leetcode.com/problems/leaf-similar-trees/submissions/1226550945/?envType=study-plan-v2&envId=leetcode-75

import . "leetcode/pkg/datastructures/treenode"

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var (
		traverse func(root *TreeNode, leaves *[]int)

		roots1, roots2 = make([]int, 0, 128), make([]int, 0, 128)
	)

	traverse = func(root *TreeNode, leaves *[]int) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			*leaves = append(*leaves, root.Val)
		}

		traverse(root.Left, leaves)
		traverse(root.Right, leaves)
	}

	traverse(root1, &roots1)
	traverse(root2, &roots2)

	if len(roots1) != len(roots2) {
		return false
	}

	for i, n := range roots1 {
		if n != roots2[i] {
			return false
		}
	}

	return true
}
