package middle

// https://leetcode.com/problems/check-knight-tour-configuration/description/

type KnightPosition struct {
	x, y       int
	stepNumber int
}

type Point2D struct {
	x, y int
}

func checkValidGrid(grid [][]int) bool {
	const (
		xInd = iota
		yInd
	)

	var (
		shifts = [8]Point2D{
			{x: -1, y: 2},
			{x: -2, y: 1},
			{x: 1, y: 2},
			{x: 2, y: 1},
			{x: -2, y: -1},
			{x: -1, y: -2},
			{x: 1, y: -2},
			{x: 2, y: -1},
		}

		boardSize = len(grid)
		queue     = make([]KnightPosition, 0, 1<<6)
		curStep   = 1
	)

	if grid[0][0] != 0 {
		return false
	}

	for _, ks := range shifts {
		queue = append(queue, KnightPosition{x: ks.x, y: ks.y, stepNumber: curStep})
	}

	for len(queue) != 0 {
		curKP := queue[len(queue)-1]

		if curKP.x < 0 || curKP.x >= boardSize {
			queue = queue[:len(queue)-1]
			continue
		}

		if curKP.y < 0 || curKP.y >= boardSize {
			queue = queue[:len(queue)-1]
			continue
		}

		if grid[curKP.x][curKP.y] != curStep {
			queue = queue[:len(queue)-1]
			continue
		}

		queue = queue[:1]
		queue[0] = curKP

		curStep++
		for _, ks := range shifts {
			queue = append(queue, KnightPosition{x: curKP.x + ks.x, y: curKP.y + ks.y, stepNumber: curStep})
		}
	}

	return curStep == boardSize*boardSize
}
