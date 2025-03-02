package middle

// https://leetcode.com/problems/score-after-flipping-matrix/description/?envType=daily-question&envId=2024-05-13

import "math"

func matrixScore(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		if grid[i][0] == 0 {
			for j := 0; j < len(grid[0]); j++ {
				if grid[i][j] == 0 {
					grid[i][j] = 1
					continue
				}
				grid[i][j] = 0
			}
		}
	}

	var (
		curColZerosNumber int
	)
out:
	for i := 0; i < len(grid[0]); i++ {
		curColZerosNumber = 0

		for j := 0; j < len(grid); j++ {
			if grid[j][i] == 0 {
				curColZerosNumber++
			}

			if curColZerosNumber > len(grid)/2 {
				for k := 0; k < len(grid); k++ {
					if grid[k][i] == 0 {
						grid[k][i] = 1
						continue
					}
					grid[k][i] = 0
				}

				continue out
			}
		}

	}

	var (
		curPow       int
		resultingSum int
	)
	for i := 0; i < len(grid); i++ {
		curPow = len(grid[0]) - 1
		for j := 0; j < len(grid[0]); j++ {
			resultingSum += grid[i][j] * int(math.Pow(2, float64(curPow)))
			curPow--
		}
	}

	return resultingSum

}
