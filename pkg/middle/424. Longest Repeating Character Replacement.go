package middle

func characterReplacement(s string, k int) int {
	var (
		usedLts = getUsedLetters(s)

		maxSequenceLen int
	)

	if k == 0 {
		return findMaxLetterSubsequenceLen(s)
	}

	for rep := range usedLts {
		maxSequenceLen = max(maxSequenceLen, findMaxSubstitutedSequenceLength(rep, k, s))
	}

	return maxSequenceLen
}

func findMaxSubstitutedSequenceLength(rep byte, k int, str string) int {
	var (
		repIdx, repCnt     = make([]int, 0, 1<<6), 0
		l, r           int = 0, 0

		maxLen int
	)

	for r < len(str) {

		if str[r] != rep {
			switch {
			case repCnt < k:
				repIdx = append(repIdx, r)
				repCnt++

			case repCnt == k:
				if curLen := r - l; curLen > maxLen {
					maxLen = curLen
				}
				l, repIdx = repIdx[0]+1, repIdx[1:]

				repIdx = append(repIdx, r)

			}

		}

		r++
	}

	if curLen := r - l; curLen > maxLen {
		maxLen = curLen
	}

	return maxLen
}

func findMaxLetterSubsequenceLen(str string) int {
	var (
		maxLen = 1

		l, r int = 0, 1
	)

	for r < len(str) {

		for r < len(str) && str[r] == str[l] {
			r++
		}

		if curLen := r - l; curLen > maxLen {
			maxLen = curLen
		}

		l, r = r, r+1

	}

	if curLen := r - l; curLen > maxLen {
		maxLen = curLen
	}

	return maxLen
}

func getUsedLetters(s string) map[byte]struct{} {
	var (
		m = make(map[byte]struct{}, 1<<5)
	)

	for i := range s {
		m[s[i]] = struct{}{}
	}

	return m
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}

// 	return b
// }
