package easy

// https://leetcode.com/problems/largest-local-values-in-a-matrix/description/?envType=daily-question&envId=2024-05-12

func largestLocal(grid [][]int) [][]int {

	var (
		result = make([][]int, len(grid)-2)
	)
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, len(result))
	}

	var (
		rowPad, colPad int

		sectorMax int
	)

	for rowPad+2 < len(grid) && colPad+2 < len(grid) {
		sectorMax = -1

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if grid[i+rowPad][j+colPad] > sectorMax {
					sectorMax = grid[i+rowPad][j+colPad]
				}
			}
		}

		result[rowPad][colPad] = sectorMax
		sectorMax = -1

		colPad++
		if colPad+2 >= len(grid) {
			colPad = 0
			rowPad++
		}
	}

	return result
}
