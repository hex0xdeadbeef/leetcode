package middle

// https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/?envType=study-plan-v2&envId=leetcode-75

import (
	. "leetcode/pkg/datastructures/treenode"
)

// type zigzagLine struct {
// 	curLength int
// 	direction int
// }

// func longestZigZag(root *TreeNode) int {

// 	const (
// 		left  = -1
// 		right = 1
// 	)

// 	var (
// 		maxZigZagLength = 0

// 		traverse      func(root *TreeNode)
// 		exploreZigzag func(root *TreeNode, zigZag zigzagLine)
// 	)

// 	traverse = func(root *TreeNode) {
// 		if root == nil {
// 			return
// 		}

// 		exploreZigzag(root, zigzagLine{direction: -1})
// 		exploreZigzag(root, zigzagLine{direction: 1})

// 		traverse(root.Left)
// 		traverse(root.Right)
// 	}

// 	exploreZigzag = func(root *TreeNode, zigZag zigzagLine) {
// 		switch zigZag.direction {
// 		case left:
// 			if canRight(root) {
// 				exploreZigzag(root.Right, zigzagLine{curLength: zigZag.curLength + 1, direction: 1})
// 			}

// 		case right:
// 			if canLeft(root) {
// 				exploreZigzag(root.Left, zigzagLine{curLength: zigZag.curLength + 1, direction: -1})
// 			}

// 		default:
// 			return
// 		}

// 		if zigZag.curLength > maxZigZagLength {
// 			maxZigZagLength = zigZag.curLength
// 		}

// 	}
// 	traverse(root)

// 	return maxZigZagLength
// }

// func canRight(root *TreeNode) bool {
// 	return root.Right != nil
// }

// func canLeft(root *TreeNode) bool {
// 	return root.Left != nil
// }

func longestZigZag(root *TreeNode) int {
	return max(checkZigZag(root.Left, true, 0), checkZigZag(root.Right, false, 0))
}

func checkZigZag(head *TreeNode, leftOrRight bool, count int) int {
	if head == nil {
		return count
	}

	if leftOrRight {
		return max(checkZigZag(head.Right, false, count+1), checkZigZag(head.Left, true, 0))
	} else {
		return max(checkZigZag(head.Left, true, count+1), checkZigZag(head.Right, false, 0))
	}
}
