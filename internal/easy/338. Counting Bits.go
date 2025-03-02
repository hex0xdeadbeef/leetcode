package easy

// https://leetcode.com/problems/counting-bits/

func countBits(n int) []int {

	const (
		sizePad = 1
	)
	var (
		bitsCounts = make([]int, 0, n+sizePad)
	)

	for i := 0; i <= n; i++ {
		bitsCounts = append(bitsCounts, popCount(i))
	}

	return bitsCounts
}

func popCount(num int) int {
	var (
		count int
	)

	for num > 0 {
		count++

		num &= (num - 1)

	}

	return count
}
