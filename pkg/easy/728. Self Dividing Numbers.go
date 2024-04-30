package easy

// https://leetcode.com/problems/self-dividing-numbers/

func selfDividingNumbers(left int, right int) []int {

	const (
		pad = 1
	)
	var (
		result = make([]int, 0, right-left+pad)
	)

	for i := left; i <= right; i++ {
		if isSelfDividingNumber(i) {
			result = append(result, i)
		}
	}

	return result

}

func isSelfDividingNumber(num int) bool {
	const (
		evenReminder = 0
		divider      = 10

		zero = 0
	)

	var (
		numCopy   = num
		lastDigit int
	)

	for numCopy > 0 {
		if lastDigit = numCopy % divider; lastDigit == zero || num%lastDigit != evenReminder {
			return false
		}
		numCopy /= divider
	}

	return true

}
