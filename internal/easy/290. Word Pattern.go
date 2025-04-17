package easy

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	fromPatternToWord := map[byte]string{}
	fromWordsToPattern := map[string]byte{}

	for i, b := range []byte(pattern) {
		_, ok1 := fromPatternToWord[b]
		if !ok1 {
			fromPatternToWord[b] = words[i]
		}

		_, ok2 := fromWordsToPattern[words[i]]
		if !ok2 {
			fromWordsToPattern[words[i]] = b
		}

		if !ok1 && !ok2 {
			continue
		}

		if fromPatternToWord[b] != words[i] {
			return false
		}

		if fromWordsToPattern[words[i]] != b {
			return false
		}
	}

	return true
}
