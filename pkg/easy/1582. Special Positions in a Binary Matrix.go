package easy

func numSpecial(mat [][]int) int {

	var (
		rowsLen, columnsLen = len(mat), len(mat[0])

		rowsCounts, columnsCount = make([]int, rowsLen), make([]int, columnsLen)
	)

	for i := 0; i < rowsLen; i++ {

		for j := 0; j < columnsLen; j++ {
			if mat[i][j] == 1 {
				rowsCounts[i]++
				columnsCount[j]++
			}
		}

	}

	var (
		count int
	)

	for i := 0; i < rowsLen; i++ {

		for j := 0; j < columnsLen; j++ {
			if mat[i][j] == 1 && rowsCounts[i] == 1 && columnsCount[j] == 1 {
				count++
			}
		}

	}

	return count
}
