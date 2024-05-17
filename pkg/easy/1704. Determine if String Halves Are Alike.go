package easy

// https://leetcode.com/problems/determine-if-string-halves-are-alike/description/

import "strings"

func halvesAreAlike(s string) bool {

	var (
		vowels      = map[byte]empty{'a': {}, 'e': {}, 'o': empty{}, 'i': empty{}, 'u': empty{}}
		balance int
	)

	s = strings.ToLower(s)

	for i := 0; i < len(s)/2; i++ {
		if _, ok := vowels[s[i]]; ok {
			balance++
		}

		if _, ok := vowels[s[len(s)-1-i]]; ok {
			balance--
		}


	}

	return balance == 0
}
