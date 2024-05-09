package easy

// https://leetcode.com/problems/minimum-string-length-after-removing-substrings/description/

func minLength(s string) int {

	const (
		sub1, sub2 = "AB", "CD"
		idxPad     = 1
		pairIdxPad = 2
	)

	var (
		stack = make([]rune, 0, len(s))
		top   string
	)

	for _, v := range s {
		stack = append(stack, v)

		if len(stack) >= 2 {
			if top = string(stack[len(stack)-pairIdxPad]) + string(stack[len(stack)-idxPad]); top == sub1 || top == sub2 {
				stack = stack[:len(stack)-pairIdxPad]
			}
		}

	}

	return len(stack)
}
