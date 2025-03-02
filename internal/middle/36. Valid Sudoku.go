package middle

// https://leetcode.com/problems/valid-sudoku/

func isValidSudoku(board [][]byte) bool {
	const (
		clusterSize                         = 9
		clusterColumnShift, clusterRowShift = 3, 3
		maxClusterColumnPad                 = 9
	)

	var (
		rows, cols = make([]map[byte]empty, 9), make([]map[byte]empty, 9)
	)

	for i := 0; i < len(board); i++ {
		rows[i], cols[i] = make(map[byte]empty, 9), make(map[byte]empty, 9)
		for j := 0; j < len(board); j++ {
			if board[i][j] != '.' {
				if _, ok := rows[i][board[i][j]]; ok {
					return false
				}
				rows[i][board[i][j]] = empty{}
			}

			if board[j][i] != '.' {

				if _, ok := cols[i][board[j][i]]; ok {
					return false
				}
				cols[i][board[j][i]] = empty{}
			}
		}
	}

	var (
		cluster  map[byte]empty
		colShift int = 0
	)

	for colShift < maxClusterColumnPad {

		for k := 0; k < len(board); k += clusterRowShift {
			cluster = make(map[byte]empty, clusterSize)

			for j := k; j < k+clusterRowShift; j++ {

				for i := colShift; i < colShift+clusterColumnShift; i++ {
					if _, ok := cluster[board[j][i]]; ok {
						return false
					}
					if board[j][i] != '.' {
						cluster[board[j][i]] = empty{}
					}

				}

			}
		}

		colShift += clusterColumnShift
	}

	return true
}
