package easy

// https://leetcode.com/problems/find-the-width-of-columns-of-a-grid/description/

import "strconv"

func findColumnWidth(grid [][]int) []int {

	var (
		res = make([]int, len(grid[0]))
	)

	for columnIdx := 0; columnIdx < len(grid[0]); columnIdx++ {

		for _, row := range grid {
			res[columnIdx] = max(res[columnIdx], len(strconv.Itoa(row[columnIdx])))
		}
	}

	return res
}
