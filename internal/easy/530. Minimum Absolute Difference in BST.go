package easy

import . "leetcode/pkg/datastructures/treenode"

func getMinimumDifference(root *TreeNode) int {

	var minAbsDiff *int = new(int)
	*minAbsDiff = 1e6

	var traverse func(*TreeNode, *int)

	var findMin func(*TreeNode, *int, int)

	traverse = func(root *TreeNode, minAbsDiff *int) {
		if root == nil {
			return
		}

		findMin(root, minAbsDiff, root.Val)
		traverse(root.Left, minAbsDiff)
		traverse(root.Right, minAbsDiff)
	}

	findMin = func(root *TreeNode, minAbsDiff *int, val int) {
		if root == nil {
			return
		}

		if root.Val != val && abs(root.Val-val) < *minAbsDiff {
			*minAbsDiff = abs(root.Val - val)
		}

		findMin(root.Left, minAbsDiff, val)
		findMin(root.Right, minAbsDiff, val)
	}

	traverse(root, minAbsDiff)

	return *minAbsDiff
}
