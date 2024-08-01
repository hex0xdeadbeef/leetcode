package easy

// https://leetcode.com/problems/maximum-value-of-a-string-in-an-array/

import (
	"strconv"
	"unicode"
)

func maximumValue(strs []string) int {
	var (
		maxVal = -1
	)

	for _, s := range strs {
		maxVal = max(maxVal, getNumeric(s))
	}

	return maxVal
}

func getNumeric(s string) int {

	for _, r := range s {
		if !unicode.IsNumber(r) {
			return len(s)
		}
	}

	num, _ := strconv.Atoi(s)

	return num
}
