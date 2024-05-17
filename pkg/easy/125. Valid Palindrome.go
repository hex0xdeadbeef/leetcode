package easy

// https://leetcode.com/problems/valid-palindrome/description/

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	var (
		l, r int = 0, len(s) - 1
	)
	s = strings.ToLower(s)

	for r > l {
		for r >= 0 && (!unicode.IsLetter(rune(s[r])) && !unicode.IsNumber(rune(s[r]))) {
			r--
		}

		for l < len(s) && (!unicode.IsLetter(rune(s[l])) && !unicode.IsNumber(rune(s[l]))) {
			l++
		}

		if r < 0 || l >= len(s) {
			break
		}

		if s[r] != s[l] {
			return false
		}

		r--
		l++
	}

	return true

}
