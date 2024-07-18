package middle

// https://leetcode.com/problems/permutation-in-string/description/

func checkInclusion(word string, str string) bool {
	var (
		window          = make(map[byte]int, 1<<5)
		wordLetterToCnt = make(map[byte]int, len(str))
	)

	if len(word) > len(str) {
		return false
	}

	for i := range word {
		wordLetterToCnt[word[i]]++
	}

	for i := 0; i < len(word); i++ {
		window[str[i]]++
	}

	for i := 0; i < len(str); i++ {
		if isMatched(wordLetterToCnt, window) {
			return true
		}

		window[str[i]]--
		if v, _ := window[str[i]]; v == 0 {
			delete(window, str[i])
		}

		if len(word)+i == len(str) {
			return false
		}
		window[str[len(word)+i]]++
	}

	return false
}

func isMatched(super, sub map[byte]int) bool {
	if len(super) != len(sub) {
		return false
	}

	for k, v1 := range super {
		v2, ok := sub[k]
		if !ok {
			return false
		}

		if v1 != v2 {
			return false
		}
	}

	return true
}
