package easy

// https://leetcode.com/problems/reverse-prefix-of-word/?envType=daily-question&envId=2024-05-01

import (
	"bytes"
)

func reversePrefix(word string, ch byte) string {
	const (
		indPad int = 1
	)
	var (
		wordBytes = []byte(word)
		chIndex   = bytes.IndexByte(wordBytes, ch)
	)

	reverseBytes(wordBytes[:chIndex+indPad])
	return string(wordBytes)
}

func reverseBytes(b []byte) {
	const (
		divider = 2
		indPad  = 1
	)

	var (
		length       = len(b)
		halfedLength = len(b) / divider
	)

	for i := 0; i < halfedLength; i++ {
		b[i], b[length-i-indPad] = b[length-i-indPad], b[i]
	}
}
