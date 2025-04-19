package easy

func validPath(n int, edges [][]int, source int, destination int) bool {
	m := map[int]map[int]struct{}{}

	for _, e := range edges {
		_, ok1 := m[e[0]]
		if !ok1 {
			m[e[0]] = map[int]struct{}{}
		}

		_, ok2 := m[e[1]]
		if !ok2 {
			m[e[1]] = map[int]struct{}{}
		}

		m[e[0]][e[1]] = struct{}{}
		m[e[1]][e[0]] = struct{}{}
	}

	visited := map[int]struct{}{}
	visited[source] = struct{}{}

	traverseGraphDFS(m, visited, source, destination)

	_, ok := visited[destination]
	return ok
}

func traverseGraphDFS(grid map[int]map[int]struct{}, visited map[int]struct{}, source, destination int) {
	visited[source] = struct{}{}

	for dest := range grid[source] {
		_, ok := visited[dest]
		if ok {
			continue
		}

		traverseGraphDFS(grid, visited, dest, dest)
	}

}
