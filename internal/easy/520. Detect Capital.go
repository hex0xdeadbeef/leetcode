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
