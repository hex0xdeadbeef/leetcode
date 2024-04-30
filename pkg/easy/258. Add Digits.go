package easy

// https://leetcode.com/problems/add-digits/

func AddDigits(num int) int {

	const (
		divider = 10
	)
	var (
		prev int
		cur  = num
	)

	for cur / divider > 0 {
		prev = 0
		for cur > 0 {
			prev += cur % divider
			cur /= divider
		}
		cur = prev
	}

	return cur
}
