package middle

// https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/?envType=study-plan-v2&envId=leetcode-75

import (
	. "leetcode/pkg/datastructures/linkedlist"
)

func pairSum(head *ListNode) int {
	var (
		elemsCopy = make([]int, 0, 8192)
		cur       = head

		maxSum int
	)

	for cur != nil {
		elemsCopy = append(elemsCopy, cur.Val)
		cur = cur.Next
	}

	for i := 0; i < len(elemsCopy)/2; i++ {
		if curTwinSum := elemsCopy[i] + elemsCopy[len(elemsCopy)-i-1]; curTwinSum > maxSum {
			maxSum = curTwinSum
		}
	}

	return maxSum
}
