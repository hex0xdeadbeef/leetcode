package easy

// https://leetcode.com/problems/number-of-senior-citizens/description/?envType=daily-question&envId=2024-08-01

import "strconv"

func countSeniors(details []string) int {

	const (
		phoneNumberSymbolsCnt = 10
		genderSymbolCnt       = 1
		leftOverallPad        = phoneNumberSymbolsCnt + genderSymbolCnt

		rightIdxPad     = 2
		rightOverallPad = leftOverallPad + rightIdxPad

		edgeAge = 60
	)
	var (
		cnt int
	)

	for _, s := range details {

		if v, _ := strconv.Atoi(s[leftOverallPad:rightOverallPad]); v > edgeAge {
			cnt++
		}

	}
	return cnt
}
