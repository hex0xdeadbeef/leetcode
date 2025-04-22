package middle

import "container/list"

func findCircleNum(isConnected [][]int) int {
	if len(isConnected) == 1 && isConnected[0][0] == 1 {
		return 1
	}

	paths := map[int]map[int]struct{}{}
	start := -1

	for i, edges := range isConnected {
		_, ok := paths[i]
		if !ok {
			paths[i] = map[int]struct{}{}
		}

		for j, isRoadExist := range edges {
			if isRoadExist != 1 {
				continue
			}

			if start == -1 {
				start = i
			}

			paths[i][j] = struct{}{}
		}
	}

	var componentsNumber int

	visited := map[int]struct{}{}
	visited[start] = struct{}{}

	q := list.New()
	q.PushBack(start)

Out:
	for len(visited) != len(paths) {

		for q.Len() != 0 {
			e := q.Front()
			node := e.Value.(int)
			visited[node] = struct{}{}
			q.Remove(e)

			for nextNode, _ := range paths[node] {
				_, ok := visited[nextNode]
				if ok {
					continue
				}

				q.PushBack(nextNode)
			}
		}

		componentsNumber++

		for startNode, _ := range paths {
			_, ok := visited[startNode]
			if ok {
				continue
			}

			q.PushBack(startNode)
			continue Out
		}
	}

	return componentsNumber
}
