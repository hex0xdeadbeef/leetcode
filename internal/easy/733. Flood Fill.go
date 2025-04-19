package easy

import "container/list"

type coordinate struct {
	h int
	v int
}

var steps = []coordinate{
	{h: 0, v: 1},
	{h: 0, v: -1},
	{h: 1, v: 0},
	{h: -1, v: 0},
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	visited := map[coordinate]struct{}{}

	q := list.New()
	q.PushBack(coordinate{h: sr, v: sc})
	visited[q.Front().Value.(coordinate)] = struct{}{}

	for q.Len() != 0 {
		e := q.Front()
		c := e.Value.(coordinate)
		visited[c] = struct{}{}
		q.Remove(e)

		prevColor := image[c.h][c.v]
		image[c.h][c.v] = color

		for _, s := range steps {
			newCoordinate := coordinate{h: c.h + s.h, v: c.v + s.v}

			_, ok := visited[newCoordinate]
			if ok {
				continue
			}

			if newCoordinate.h < 0 || newCoordinate.h >= len(image) {
				continue
			}

			if newCoordinate.v < 0 || newCoordinate.v >= len(image[0]) {
				continue
			}

			if image[newCoordinate.h][newCoordinate.v] != prevColor {
				continue
			}

			q.PushBack(newCoordinate)
		}
	}

	return image
}
