package middle

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

func lengthOfLongestSubstring(s string) int {
	var (
		lettersIndices = make(map[byte]int, 1<<5)
		subLetters     = make(map[byte]struct{}, 1<<4)

		l, r = 0, 1

		maxLen int
	)

	if s == "" {
		return 0
	}

	subLetters[s[l]], lettersIndices[s[l]] = struct{}{}, l

	for r < len(s) {
		switch _, ok := subLetters[s[r]]; ok {

		case true:
			if curLen := len(subLetters); curLen > maxLen {
				maxLen = curLen
			}

			removeInvalidLetters(l, lettersIndices[s[r]], s, subLetters)
			l, lettersIndices[s[r]] = lettersIndices[s[r]]+1, r

		default:
			subLetters[s[r]], lettersIndices[s[r]] = struct{}{}, r
		}

		r++
	}

	if curLen := len(subLetters); curLen > maxLen {
		maxLen = curLen
	}

	return maxLen
}

func removeInvalidLetters(l, r int, str string, subLetters map[byte]struct{}) {
	for ; l < r; l++ {
		delete(subLetters, str[l])
	}
}

func lengthOfLongestSubstringRepeat(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	var l, maxLen int
	m := make(map[byte]int, len(s))

	for r := 0; r < len(s); r++ {
		if prevIdx, ok := m[s[r]]; ok && prevIdx >= l {
			l = prevIdx + 1
		}

		m[s[r]] = r
		if curLen := r - l + 1; curLen > maxLen {
			maxLen = curLen
		}
	}

	return maxLen
}
