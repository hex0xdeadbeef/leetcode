package easy

func reverseWords(s string) string {
	b := []byte(s)

	var l, r int

	for l < len(b) {
		for ; l < len(b) && b[l] == ' '; l++ {
		}

		if l == len(b) {
			break
		}

		r = l
		for ; r < len(b) && b[r] != ' '; r++ {
		}

		for i, j := l, r-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}

		l = r
	}

	return string(b)
}
