package middle

import (
	"container/list"
	"fmt"
	. "leetcode/pkg/datastructures/graphnode"
)

// https://leetcode.com/problems/number-of-islands/description/

var (
	kingShifts = [4][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
)

func numIslands(grid [][]byte) int {
	var (
		islandsCnt int
	)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				islandsCnt++
				grid[i][j] = '0'
				islandsDFS(i, j, grid)
			}
		}
	}

	return islandsCnt
}

func islandsDFS(rowIdx, colIdx int, grid [][]byte) {
	for _, shift := range kingShifts {
		y, x := rowIdx+shift[0], colIdx+shift[1]

		if y < 0 || y >= len(grid) {
			continue
		}

		if x < 0 || x >= len(grid[0]) {
			continue
		}

		if grid[y][x] == '1' {
			grid[y][x] = '0'
			islandsDFS(y, x, grid)
		}
	}
}

// https://codeforces.com/group/dnpUeCC2dg/contest/383829/problem/D

func labyrinth(grid [][]byte, k int) {
	var (
		dotsCnt         int
		overallStepsCnt int
	)

	for i, row := range grid {
		for j, char := range row {
			if char == '.' {
				dotsCnt++
				grid[i][j] = 'X'
			}
		}
	}
	overallStepsCnt = dotsCnt - k

Out:
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'X' {
				grid[i][j] = '.'
				overallStepsCnt--
				labyrinthDFS(grid, i, j, &overallStepsCnt)
				break Out
			}
		}
	}

}

func labyrinthDFS(grid [][]byte, rowIdx, colIdx int, overallStepsCntPtr *int) {
	for _, shift := range kingShifts {
		y, x := rowIdx+shift[0], colIdx+shift[1]

		if *overallStepsCntPtr == 0 {
			return
		}

		if y < 0 || y >= len(grid) {
			continue
		}

		if x < 0 || x >= len(grid[0]) {
			continue
		}

		if grid[y][x] == 'X' {
			*overallStepsCntPtr--

			grid[y][x] = '.'
			labyrinthDFS(grid, y, x, overallStepsCntPtr)
		}
	}
}

func numIslandsNew(grid [][]byte) int {
	var islandsCount int

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] != '1' {
				continue
			}

			grid = traverseIslandBFS(grid, Coordinate{Row: i, Col: j})
			islandsCount++
		}
	}

	fmt.Println(grid)

	return islandsCount
}

func traverseIslandBFS(grid [][]byte, c Coordinate) [][]byte {

	q := list.New()
	q.PushBack(c)

	grid[c.Row][c.Col] = '0'

	for q.Len() != 0 {
		e := q.Front()
		c := e.Value.(Coordinate)
		q.Remove(e)

		for _, s := range Steps {

			newC := Coordinate{Row: c.Row + s.Row, Col: c.Col + s.Col}

			if newC.Row < 0 || newC.Row >= len(grid) || newC.Col < 0 || newC.Col >= len(grid[0]) {
				continue
			}

			if grid[newC.Row][newC.Col] != '1' {
				continue
			}

			grid[newC.Row][newC.Col] = '0'
			q.PushBack(newC)
		}
	}

	return grid
}
