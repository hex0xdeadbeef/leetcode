package improve

import "fmt"

// https://leetcode.com/problems/find-the-town-judge/

func FindJudge(n int, trust [][]int) int {

	if n == 1 {
		return 1
	}

	if len(trust) == 0 {
		return -1
	}

	var (
		seen          = make(map[int][][]int, n)
		probableJudge int
	)

	for _, e := range trust {
		if _, ok := seen[e[0]]; !ok {
			seen[e[0]] = make([][]int, 0, n)
		}

		seen[e[0]] = append(seen[e[0]], e)
	}

	if len(seen) != n-1 {
		return -1
	}

	for k, v := range seen {
		fmt.Printf("Key: %d -> Values: %v\n", k, v)
	}

	for i := 1; i <= n; i++ {
		if _, ok := seen[i]; !ok {
			probableJudge = i
			break
		}
	}

	for _, m := range seen {
		for _, v := range m {
			if v[1] == probableJudge {
				goto judgeabsence
			}
		}

		return -1

	judgeabsence:
		continue
	}

	return probableJudge

}
