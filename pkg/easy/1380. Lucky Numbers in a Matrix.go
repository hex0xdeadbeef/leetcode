package easy

// https://leetcode.com/problems/lucky-numbers-in-a-matrix/

func luckyNumbers(matrix [][]int) []int {

	var (
		rowsMins    = make([]int, len(matrix))
		columnsMaxs = make([]int, len(matrix[0]))

		result = make([]int, 0, len(matrix))
	)

	for i := 0; i < len(matrix); i++ {
		var curMin int = 10e6

		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] < curMin {
				curMin = matrix[i][j]
			}
		}

		rowsMins[i] = curMin
	}

	for i := 0; i < len(matrix[0]); i++ {
		var curMax int = 0

		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] > curMax {
				curMax = matrix[j][i]
			}
		}

		columnsMaxs[i] = curMax
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == rowsMins[i] && matrix[i][j] == columnsMaxs[j] {
				result = append(result, matrix[i][j])
			}
		}
	}

	return result
}
