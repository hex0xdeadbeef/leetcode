package easy

// https://leetcode.com/problems/number-of-common-factors/description/

func commonFactors(a int, b int) int {
	var (
		dividersCnt int

		maxNum, minNum = max(a, b), min(a, b)
	)

	for i := 1; i <= minNum; i++ {
		if maxNum%i == 0 && minNum%i == 0 {
			dividersCnt++
		}
	}

	return dividersCnt

}
