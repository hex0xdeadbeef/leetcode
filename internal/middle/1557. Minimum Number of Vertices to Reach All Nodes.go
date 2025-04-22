package middle

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	graph := map[int]map[int]struct{}{}

	for _, e := range edges {
		_, ok := graph[e[1]]
		if !ok {
			graph[e[1]] = map[int]struct{}{}
		}

		graph[e[1]][e[0]] = struct{}{}
	}

	res := []int{}

	for i := range n {
		_, ok := graph[i]
		if !ok {
			res = append(res, i)
		}
	}

	return res
}
