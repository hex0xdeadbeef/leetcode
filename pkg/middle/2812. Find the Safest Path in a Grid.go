package middle

import (
	"container/heap"
	"container/list"
	"fmt"
)

// https://leetcode.com/problems/find-the-safest-path-in-a-grid/description/?envType=daily-question&envId=2024-05-15

// type CoordinatesPair struct {
// 	x, y int
// }

// var (
// 	shifts = []CoordinatesPair{
// 		{0, 1},
// 		{0, -1},
// 		{1, 0},
// 		{-1, 0},
// 	}
// )

// type empty struct{}

type DistanceCoord struct {
	dist int
	x, y int
}

type DistancesCoords []DistanceCoord

func (a *DistancesCoords) Len() int           { return len(*a) }
func (a *DistancesCoords) Swap(i, j int)      { (*a)[i], (*a)[j] = (*a)[j], (*a)[i] }
func (a *DistancesCoords) Less(i, j int) bool { return (*a)[i].dist > (*a)[j].dist }
func (a *DistancesCoords) Pop() interface{} {
	old := *a
	n := len(old)
	x := old[n-1]
	*a = old[0 : n-1]
	return x
}
func (a *DistancesCoords) Push(x interface{}) {
	*a = append(*a, x.(DistanceCoord))
}

func MaximumSafenessFactor(grid [][]int) int {
	const (
		inf = 1e9
	)

	var (
		thieves = list.New()
	)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				thieves.PushFront(CoordinatesPair{x: i, y: j})
				grid[i][j] = 0
				continue
			}
			grid[i][j] = inf
		}
	}

	bfs(grid, thieves)
	return dijkstra(grid)
}

func bfs(grid [][]int, queue *list.List) {

	var (
		curElem, newCoord CoordinatesPair
	)

	for queue.Len() != 0 {
		curElem = queue.Front().Value.(CoordinatesPair)
		queue.Remove(queue.Front())

		for _, v := range shifts {

			newCoord = CoordinatesPair{x: curElem.x + v.x, y: curElem.y + v.y}

			if newCoord.x < 0 || newCoord.x >= len(grid) {
				continue
			}

			if newCoord.y < 0 || newCoord.y >= len(grid[0]) {
				continue
			}

			if newCoordVal := grid[curElem.x][curElem.y] + 1; newCoordVal < grid[newCoord.x][newCoord.y] {
				grid[newCoord.x][newCoord.y] = newCoordVal
				queue.PushBack(newCoord)
			}

		}
	}

}

func dijkstra(grid [][]int) int {
	for _, v := range grid {
		fmt.Println(v)
	}
	var (
		h = &DistancesCoords{DistanceCoord{dist: grid[0][0], x: 0, y: 0}}

		curElem, newElem DistanceCoord

		visited = make(map[CoordinatesPair]empty, len(grid)*len(grid[0]))
	)
	heap.Init(h)

	for h.Len() != 0 {
		curElem = heap.Pop(h).(DistanceCoord)
		visited[CoordinatesPair{x: curElem.x, y: curElem.y}] = empty{}

		if curElem.x == len(grid)-1 && curElem.y == len(grid[0])-1 {
			return curElem.dist
		}

		for _, v := range shifts {
			newElem = DistanceCoord{x: curElem.x + v.x, y: curElem.y + v.y}

			if newElem.x < 0 || newElem.x >= len(grid) {
				continue
			}

			if newElem.y < 0 || newElem.y >= len(grid[0]) {
				continue
			}

			if _, ok := visited[CoordinatesPair{newElem.x, newElem.y}]; ok {
				continue
			}

			visited[CoordinatesPair{newElem.x, newElem.y}] = empty{}

			if curElem.dist < grid[newElem.x][newElem.y] {
				newElem.dist = curElem.dist
			} else {
				newElem.dist = grid[newElem.x][newElem.y]
			}

			heap.Push(h, newElem)
		}
	}

	panic("UNREACHABLE")
}
