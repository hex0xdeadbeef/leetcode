package middle

import "slices"

func allPathsSourceTarget(graph [][]int) [][]int {
	var res [][]int

	curPath := []int{}

	curPath = append(curPath, 0)

	allPathsSourceTargetDFS(0, graph, curPath, &res)

	return res
}

func allPathsSourceTargetDFS(start int, graph [][]int, curPath []int, path *[][]int) bool {

	if start == len(graph)-1 {
		return true
	}

	for _, next := range graph[start] {

		curPath = append(curPath, next)

		if allPathsSourceTargetDFS(next, graph, curPath, path) {
			*path = append(*path, slices.Clone(curPath))
		}

		curPath = curPath[:len(curPath)-1]
	}

	return false
}
