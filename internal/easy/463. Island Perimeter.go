package easy

import "container/list"

func islandPerimeter(grid [][]int) int {
	var start coordinate

Out:
	for i, r := range grid {
		for j, c := range r {
			if c == 1 {
				start.h, start.v = i, j
				break Out
			}
		}
	}

	var perimeter int

	visited := map[coordinate]struct{}{}
	visited[start] = struct{}{}

	q := list.New()
	q.PushBack(start)

	for q.Len() != 0 {
		e := q.Front()
		c := e.Value.(coordinate)
		visited[c] = struct{}{}
		q.Remove(e)

		for _, s := range steps {
			newCoordinate := coordinate{h: c.h + s.h, v: c.v + s.v}

			if newCoordinate.h < 0 || newCoordinate.h >= len(grid) || newCoordinate.v < 0 || newCoordinate.v >= len(grid[0]) ||
				grid[newCoordinate.h][newCoordinate.v] == 0 {
				perimeter++
				continue
			}

			_, ok := visited[newCoordinate]
			if ok {
				continue
			}

			if grid[newCoordinate.h][newCoordinate.v] != 1 {
				continue
			}

			q.PushBack(newCoordinate)
		}
	}

	return perimeter

}
