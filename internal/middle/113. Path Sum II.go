package middle

import . "leetcode/pkg/datastructures/treenode"

func pathSumTwo(root *TreeNode, targetSum int) [][]int {
	var traverse func(root *TreeNode, nums []int, curSum int)

	var res [][]int
	traverse = func(root *TreeNode, nums []int, curSum int) {
		if root == nil {
			return
		}

		curSum += root.Val
		nums = append(nums, root.Val)

		if root.Left == nil && root.Right == nil {
			if curSum == targetSum {
				res = append(res, append([]int(nil), nums...))
			}

			return
		}

		traverse(root.Left, nums, curSum)
		traverse(root.Right, nums, curSum)
	}

	traverse(root, nil, 0)

	return res
}
