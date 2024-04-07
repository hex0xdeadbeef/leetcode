package easy

// https://leetcode.com/problems/reverse-vowels-of-a-string/description/?envType=study-plan-v2&envId=leetcode-75

type vowel struct {
	character rune
	origIndex int
}

type empty struct{}

func reverseVowels(s string) string {
	var (
		strBytes = []byte(s)

		vowelsMap = map[rune]empty{
			'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {},
			'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {},
		}
		vowels = make([]vowel, 0, len(s))
		l      int
	)

	for i, r := range s {
		if _, ok := vowelsMap[r]; ok {
			vowels = append(vowels, vowel{character: r, origIndex: i})
		}
	}

	l = len(vowels)

	for i := 0; i < len(vowels)/2; i++ {
		strBytes[vowels[l-i-1].origIndex], strBytes[vowels[i].origIndex] = strBytes[vowels[i].origIndex], strBytes[vowels[l-i-1].origIndex]
	}

	return string(strBytes)
}
