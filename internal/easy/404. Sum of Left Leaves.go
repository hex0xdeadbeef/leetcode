package easy

// https://leetcode.com/problems/sum-of-left-leaves/

import . "leetcode/pkg/datastructures/treenode"

func sumOfLeftLeaves(root *TreeNode) int {

	var (
		leftSum int

		traverse func(root *TreeNode, isLeft bool)
	)

	traverse = func(root *TreeNode, isLeft bool) {
		if root == nil {
			return
		}

		if isLeft && (root.Left == nil && root.Right == nil) {
			leftSum += root.Val
			return
		}

		traverse(root.Left, true)
		traverse(root.Right, false)
	}

	traverse(root.Left, true)
	traverse(root.Right, false)

	return leftSum

}

func sumOfLeftLeavesNew(root *TreeNode) int {
	var sum int

	var traverse func(*TreeNode, bool)

	traverse = func(root *TreeNode, isLeft bool) {
		if root == nil {
			return
		}

		if isLeft && root.Left == nil && root.Right == nil {
			sum += root.Val
		}

		traverse(root.Left, true)
		traverse(root.Left, false)
	}

	traverse(root, false)

	return sum

}
