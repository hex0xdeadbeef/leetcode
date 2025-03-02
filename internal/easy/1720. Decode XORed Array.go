package easy

// https://leetcode.com/problems/decode-xored-array/description/

func decode(encoded []int, first int) []int {
	const (
		sizePad = 1
	)
	var (
		decoded = make([]int, 0, len(encoded)+sizePad)
	)
	decoded = append(decoded, first)

	for i, target := range encoded {
		decoded = append(decoded, target^decoded[i])

	}

	return decoded
}
