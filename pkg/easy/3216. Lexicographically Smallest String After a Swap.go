package easy

// https://leetcode.com/problems/lexicographically-smallest-string-after-a-swap/description/

func getSmallestString(s string) string {
	const (
		idxPad = 1
	)

	var (
		baseBytes = []byte(s)
	)

	for i := 0; i < len(baseBytes)-1; i++ {
		cur, next := baseBytes[i], baseBytes[i+idxPad]

		if isSameParity(cur, next) {
			if next < cur {
				baseBytes[i], baseBytes[i+idxPad] = baseBytes[i+idxPad], baseBytes[i]
				return string(baseBytes)
			}

		}
	}

	return s
}

func isSameParity(a, b byte) bool {
	const (
		pad = '0'
	)

	return (a-pad)%2 == (b-pad)%2
}
