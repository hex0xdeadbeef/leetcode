package easy

// https://leetcode.com/problems/number-of-changing-keys/

import "strings"

func countKeyChanges(s string) int {

	const (
		idxPad = 1
	)

	var (
		loweredStr = strings.ToLower(s)

		count int
	)

	for i := 0; i < len(s)-idxPad; i++ {
		if loweredStr[i] != loweredStr[i+idxPad] {
			count++
		}
	}

	return count

}
