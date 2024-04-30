package easy

// https://leetcode.com/problems/find-nearest-point-that-has-the-same-x-or-y-coordinate/

func nearestValidPoint(x int, y int, points [][]int) int {
	const (
		xInd = 0
		yInd = 1

		inf int = 10e5
	)

	var (
		curManhattanDistance = inf
		ind                  = -1
	)

	for i, v := range points {
		if x == v[xInd] || y == v[yInd] {
			if newManhattanDistance := manhattanDistance(x, y, v[xInd], v[yInd]); newManhattanDistance < curManhattanDistance {
				curManhattanDistance = newManhattanDistance
				ind = i
			}
		}
	}

	if curManhattanDistance == inf {
		return -1
	}

	return ind
}

func manhattanDistance(x1, y1, x2, y2 int) int {
	return subtractionAbs(x1, x2) + subtractionAbs(y1, y2)
}

func subtractionAbs(a, b int) int {
	if a-b < 0 {
		return -(a - b)
	}

	return a - b
}
