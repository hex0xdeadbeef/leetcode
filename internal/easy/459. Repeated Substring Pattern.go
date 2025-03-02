package easy

// https://leetcode.com/problems/repeated-substring-pattern/description/

import (
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	return strings.Contains((s+s)[1:2*len(s)-1], s)
}