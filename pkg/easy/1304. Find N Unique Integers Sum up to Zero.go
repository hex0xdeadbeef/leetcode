package easy

// https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/description/

func sumZero(n int) []int {

	const (
		pad = 1

		two          = 2
		oneRemainder = 1
	)

	var (
		result = make([]int, n)

		idx    int
		curNum int = 1
	)

	if n%two == oneRemainder {
		idx = 1
	} else {
		idx = 0
	}

	for ; idx < n-1; idx += 2 {
		result[idx], result[idx+pad] = -curNum, curNum

		curNum++
	}

	return result
}
