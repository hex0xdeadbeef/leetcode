package easy

func findCenter(edges [][]int) int {
	nodeToDestinations := map[int]map[int]struct{}{}

	for i := range edges {
		_, ok1 := nodeToDestinations[edges[i][0]]
		if !ok1 {
			nodeToDestinations[edges[i][0]] = map[int]struct{}{}
		}
		nodeToDestinations[edges[i][0]][edges[i][1]] = struct{}{}

		_, ok2 := nodeToDestinations[edges[i][1]]
		if !ok2 {
			nodeToDestinations[edges[i][1]] = map[int]struct{}{}
		}
		nodeToDestinations[edges[i][1]][edges[i][0]] = struct{}{}
	}

	var res int
	for k, destinations := range nodeToDestinations {
		if len(destinations) == len(nodeToDestinations)-1 {
			res = k
			break
		}
	}

	return res
}
