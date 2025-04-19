package easy

import . "leetcode/pkg/datastructures/treenode"

func findMode(root *TreeNode) []int {
	m := map[int]int{}

	var traverse func(*TreeNode, map[int]int)

	traverse = func(root *TreeNode, m map[int]int) {
		if root == nil {
			return
		}

		m[root.Val]++

		traverse(root.Left, m)
		traverse(root.Right, m)
	}

	traverse(root, m)

	var (
		modes    []int
		maxCount int
	)

	for _, v := range m {
		if v > maxCount {
			maxCount = v
		}
	}

	for k, v := range m {
		if v == maxCount {
			modes = append(modes, k)
		}
	}

	return modes
}
