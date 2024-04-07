package middle

// https://leetcode.com/problems/equal-row-and-column-pairs/description/?envType=study-plan-v2&envId=leetcode-75s

func equalPairs(grid [][]int) int {
	var (
		rows         = make(map[string]int, len(grid))
		matrixVector = make([]rune, len(grid))

		count int
	)

	for i := 0; i < len(grid); i++ {

		for j := 0; j < len(grid); j++ {
			matrixVector[j] = rune(grid[i][j])
		}

		rows[string(matrixVector)]++

	}

	for i := 0; i < len(grid); i++ {

		for j := 0; j < len(grid); j++ {
			matrixVector[j] = rune(grid[j][i])
		}

		if num, ok := rows[string(matrixVector)]; ok {
			count += num
		}
	}

	return count

}
