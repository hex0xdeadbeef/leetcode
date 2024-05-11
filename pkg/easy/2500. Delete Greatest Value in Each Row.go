package easy

// https://leetcode.com/problems/delete-greatest-value-in-each-row/description/

import "sort"

func DeleteGreatestValue(grid [][]int) int {
	var (
		curMaxOperand int = -1

		sum int
	)

	for i := 0; i < len(grid); i++ {
		sort.Ints(grid[i])
	}

	for len(grid[0]) != 0 {
		for i := 0; i < len(grid); i++ {
			curMaxOperand = max(grid[i][len(grid[i])-1], curMaxOperand)
			grid[i] = grid[i][:len(grid[i])-1]
		}

		sum += curMaxOperand
		curMaxOperand = -1
	}

	return sum
}
