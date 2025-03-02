package easy

// https://leetcode.com/problems/largest-positive-integer-that-exists-with-its-negative/description/

// type empty struct {}

func findMaxK(nums []int) int {
	const (
		negativeInf = -1e5
	)
	var (
		seenNegative, seenPositive = make(map[int]empty, len(nums)), make(map[int]empty, len(nums))
	)

	for _, v := range nums {
		switch v > 0 {
		case true:
			seenPositive[v] = empty{}
		default:
			seenNegative[v] = empty{}
		}
	}

	var curMax int = negativeInf

	for k, _ := range seenPositive {
		if k < curMax {
			continue
		}

		if _, ok := seenNegative[-k]; !ok {
			continue
		}

		curMax = k
	}

	if curMax == negativeInf {
		return -1
	}

	return curMax
}
