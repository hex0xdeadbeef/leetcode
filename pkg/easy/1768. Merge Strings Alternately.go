package easy

import (
	"bytes"
	"fmt"
)

// https://leetcode.com/problems/merge-strings-alternately/?envType=study-plan-v2&envId=leetcode-75

func MergeAlternately(word1 string, word2 string) string {

	const (
		zeroInd = 0
	)
	var (
		buf = *bytes.NewBuffer(make([]byte, 0, len(word1)+len(word2)))

		p1, p2 = zeroInd, zeroInd
	)

	for p1 < len(word1) && p2 < len(word2) {
		buf.Write([]byte{word1[p1], word2[p2]})
		p1++
		p2++
	}

	switch {
	case p1 != len(word1):
		buf.WriteString(word1[p1:])
	case p2 != len(word2):
		buf.WriteString(word2[p2:])
	}

	fmt.Println(buf.String())

	return buf.String()

}
