package easy

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/binary-gap/

func binaryGap(n int) int {
	const (
		oneByte, zeroByte = '1', '0'
	)

	var (
		binNumDividers = strings.Split(strings.Trim(strconv.FormatInt(int64(n), 2), "0"), "1")
		maxDist        int
	)

	if binNumDividers[0] == "" {
		binNumDividers = binNumDividers[1:]
	}

	if binNumDividers[len(binNumDividers)-1] == "" {
		binNumDividers = binNumDividers[:len(binNumDividers)-1]
	}

	for _, v := range binNumDividers {
		if curDist := len(v) + 1; curDist > maxDist {
			maxDist = curDist
		}
	}

	return maxDist
}
