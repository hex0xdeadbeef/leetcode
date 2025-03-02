package easy

// https://leetcode.com/problems/replace-all-s-to-avoid-consecutive-repeating-characters/description/

import (
	"math/rand"
)

func modifyString(s string) string {
	const (
		sizePad  = 2
		question = '?'

		idxPad = 1
	)

	var (
		strBytes = make([]byte, len(s)+sizePad)
	)
	strBytes[0], strBytes[len(strBytes)-1] = '*', '*'
	copy(strBytes[1:], []byte(s))

	for i := 1; i < len(strBytes); i++ {
		if strBytes[i] == question {
			strBytes[i] = genLetter(strBytes[i-idxPad], strBytes[i+idxPad])
		}
	}

	return string(strBytes[1 : len(strBytes)-1])
}

func genLetter(left, right byte) byte {
	const (
		asciiPad = 97
	)

	var (
		randNum = rand.Intn(26) + asciiPad
	)

	for left == byte(randNum) || right == byte(randNum) {
		randNum = rand.Intn(26) + asciiPad
	}

	return byte(randNum)
}
