package middle

// https://leetcode.com/problems/compare-version-numbers/description/?envType=daily-question&envId=2024-05-03

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	const (
		dot string = "."
	)

	var (
		version1PartsStr, version2PartsStr []string = strings.Split(version1, "."), strings.Split(version2, ".")

		resultingLen     = max(len(version1PartsStr), len(version2PartsStr))
		version1PartsNum = convertSliceFromStringsToNums(version1PartsStr, resultingLen)
		version2PartsNum = convertSliceFromStringsToNums(version2PartsStr, resultingLen)
	)

	for i := 0; i < resultingLen; i++ {
		if version1PartsNum[i] < version2PartsNum[i] {
			return -1
		}

		if version1PartsNum[i] > version2PartsNum[i] {
			return 1
		}
	}

	return 0
}

func convertSliceFromStringsToNums(s []string, resultingLength int) []int {

	var (
		result = make([]int, resultingLength)
		idx    int
	)

	for ; idx < len(s); idx++ {
		result[idx], _ = strconv.Atoi(s[idx])
	}

	if idx < resultingLength {
		for idx < resultingLength {
			result[idx] = 0
			idx++
		}
	}

	return result
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}

// 	return b
// }
