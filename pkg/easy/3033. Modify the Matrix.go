package easy

func modifiedMatrix(matrix [][]int) [][]int {
	type Corrdinates struct {
		x, y int
	}

	const (
		sizeDivider = 2
		targetVal   = -1
	)

	var (
		targetElemsCoordinates = make([]Corrdinates, 0, len(matrix)*len(matrix[0])/sizeDivider)
		columnsMaxs            = make([]int, len(matrix[0]))

		curColumnMax int
	)

	for i := 0; i < len(matrix[0]); i++ {
		curColumnMax = -1e5

		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] > curColumnMax {
				curColumnMax = matrix[j][i]
			}

			if matrix[j][i] == targetVal {
				targetElemsCoordinates = append(targetElemsCoordinates, Corrdinates{x: j, y: i})
			}
		}

		columnsMaxs[i] = curColumnMax
	}

	for _, v := range targetElemsCoordinates {
		matrix[v.x][v.y] = columnsMaxs[v.y]
	}

	return matrix
}
