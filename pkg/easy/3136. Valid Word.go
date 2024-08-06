package easy

// https://leetcode.com/problems/valid-word/

import (
	"bytes"
	"unicode"
)

// https://leetcode.com/problems/valid-word/description/

func IsValid(word string) bool {

	var (
		isLetterIncluded, isVowelIncluded, isConsonantIncluded bool
	)

	if len(word) < 3 {
		return false
	}

	for _, b := range word {
		if unicode.IsNumber(rune(b)) {
			continue
		}

		if unicode.IsLetter(rune(b)) {
			isLetterIncluded = true

			b := bytes.ToLower([]byte{byte(b)})[0]
			if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' {
				isVowelIncluded = true
				continue
			}

			isConsonantIncluded = true
			continue
		}

		return false
	}

	return isLetterIncluded && isVowelIncluded && isConsonantIncluded
}
