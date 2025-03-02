package middle

import "strings"

// https://leetcode.com/problems/reverse-words-in-a-string/?envType=study-plan-v2&envId=leetcode-75

func reverseWords(s string) string {

	var (
		words = strings.Fields(s)
		l     = len(words)
	)

	for i := 0; i < l/2; i++ {
		words[i], words[l-i-1] = words[l-i-1], words[i]
	}

	return strings.Join(words, " ")
}
