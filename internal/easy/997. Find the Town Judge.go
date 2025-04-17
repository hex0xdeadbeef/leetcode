package easy

func findJudge(n int, trust [][]int) int {
	if len(trust) == 0 {
		return -1
	}

	trustRelations := make(map[int]map[int]struct{})

	for _, t := range trust {
		_, ok := trustRelations[t[0]]
		if !ok {
			trustRelations[t[0]] = map[int]struct{}{}
		}
		trustRelations[t[0]][t[1]] = struct{}{}
	}

	var trustedCount int
	for i := 1; i <= n; i++ {
		trustedCount = 0

		for _, t := range trustRelations {
			_, ok := t[i]
			if !ok {
				continue
			}

			trustedCount++
		}

		_, ok := trustRelations[i]
		if !ok && trustedCount == n-1 {
			return i
		}
	}

	return -1
}
