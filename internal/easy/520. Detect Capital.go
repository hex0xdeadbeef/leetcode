package easy

// https://leetcode.com/problems/detect-capital/description/

import (
	"strings"
	"unicode"
)

func detectCapitalUse(word string) bool {

	switch {
	case strings.ToUpper(word) == word:
		return true
	case strings.ToLower(word) == word:
		return true
	case string(unicode.ToUpper(rune(word[0])))+strings.ToLower(word[1:]) == word:
		return true
	}

	return false
}

func detectCapitalUseNew(word string) bool {
	var capitalsCount int

	for _, b := range []byte(word) {
		if b >= 'a' && b <= 'z' {
			continue
		}

		if b >= 'A' && b <= 'Z' {
			capitalsCount++
		}
	}

	if (capitalsCount == 1 && []byte(word)[0] >= 'A' && []byte(word)[0] <= 'Z') || capitalsCount == 0 || capitalsCount == len(word) {
		return true
	}

	return false
}
