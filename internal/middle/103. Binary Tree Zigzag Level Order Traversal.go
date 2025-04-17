package middle

import . "leetcode/pkg/datastructures/treenode"

func zigzagLevelOrder(root *TreeNode) [][]int {
	m := make(map[int][]int)

	var traverse func(*TreeNode, int)

	traverse = func(root *TreeNode, d int) {
		if root == nil {
			return
		}

		d++

		m[d] = append(m[d], root.Val)
		traverse(root.Left, d)
		traverse(root.Right, d)
	}

	traverse(root, -1)

	for i := range len(m) {
		if i%2 == 1 {
			for j := 0; j < len(m[j])/2; j++ {
				m[i][j], m[i][len(m[i])-1-i] = m[i][len(m[i])-1-i], m[i][j]
			}
		}
	}

	res := make([][]int, len(m))
	for i := range len(m) {
		res[i] = m[i]
	}

	return res
}
