package improve

import "math"

// https://leetcode.com/problems/diameter-of-binary-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var (
		visit func(tn *TreeNode, curAdd int) int

		maxCurSum int
	)

	visit = func(tn *TreeNode, curAdd int) int {
		if tn == nil {
			return curAdd
		}

		curAdd++

		left := visit(tn.Left, curAdd)
		right := visit(tn.Right, curAdd)

		if curSum := (left + right) - 2*curAdd; curSum > maxCurSum {
			maxCurSum = curSum
		}

		return int(math.Max(float64(left), float64(right)))

	}

	overallSum := visit(root.Left, 0) + visit(root.Right, 0)

	return int(math.Max(float64(maxCurSum), float64(overallSum)))
}
