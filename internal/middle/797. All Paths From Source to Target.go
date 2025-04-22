package middle

func allPathsSourceTarget(graph [][]int) [][]int {
	var res [][]int
	allPathsSourceTargetDFS(0, graph, make([]int, 1), &res)
	return res
}

func allPathsSourceTargetDFS(start int, graph [][]int, curPath []int, path *[][]int) bool {

	if start == len(graph)-1 {
		return true
	}

	for _, next := range graph[start] {
		curPath = append(curPath, next)
		if allPathsSourceTargetDFS(next, graph, curPath, path) {
			*path = append(*path, append([]int(nil), curPath...))
		}
		curPath = curPath[:len(curPath)-1]
	}

	return false
}
