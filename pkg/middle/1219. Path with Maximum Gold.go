package middle

// https://leetcode.com/problems/path-with-maximum-gold/description/?envType=daily-question&envId=2024-05-14

type CoordinatesPair struct {
	x, y int
}

var shifts = []CoordinatesPair{
	{x: -1, y: 0},
	{x: 1, y: 0},
	{x: 0, y: -1},
	{x: 0, y: 1},
}

func getMaximumGold(grid [][]int) int {
	var (
		res  int
		seen = make(map[CoordinatesPair]bool, len(grid)*len(grid[0]))
	)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 0 {
				res = max(res, gatherGold(CoordinatesPair{x: i, y: j}, grid, seen))
			}
		}
	}

	return res
}

func gatherGold(start CoordinatesPair, grid [][]int, seen map[CoordinatesPair]bool) int {

	seen[start] = true
	var res int

	for _, s := range shifts {
		newCoordinate := CoordinatesPair{x: start.x + s.x, y: start.y + s.y}

		if newCoordinate.x < 0 || newCoordinate.x >= len(grid) {
			continue
		}

		if newCoordinate.y < 0 || newCoordinate.y >= len(grid[0]) {
			continue
		}

		if grid[newCoordinate.x][newCoordinate.y] == 0 {
			continue
		}

		if status, _ := seen[newCoordinate]; status {
			continue
		}

		res = max(res, gatherGold(newCoordinate, grid, seen))
	}

	seen[start] = false

	return res + grid[start.x][start.y]
}
