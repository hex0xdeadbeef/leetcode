package easy

// https://leetcode.com/problems/check-if-all-the-integers-in-a-range-are-covered/

// type empty struct{}

func isCovered(ranges [][]int, left int, right int) bool {

	const (
		sizePad = 1

		targetLength = 0

		lIdx, rIdx = 0, 1
	)

	var (
		numbers = make(map[int]empty, right-left+sizePad)
	)

	for i := left; i <= right; i++ {
		numbers[i] = empty{}
	}

	for _, interval := range ranges {

		for i := interval[lIdx]; i <= interval[rIdx]; i++ {
			delete(numbers, i)
		}
	}

	return len(numbers) == targetLength
}
