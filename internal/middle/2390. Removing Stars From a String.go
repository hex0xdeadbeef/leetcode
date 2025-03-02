package middle

// https://leetcode.com/problems/removing-stars-from-a-string/?envType=study-plan-v2&envId=leetcode-75

import "fmt"

func removeStars(s string) string {

	const (
		asterisk rune = '*'
	)

	var (
		sByte = make([]byte, len(s))
		pos   = -1
	)

	for _, char := range s {
		if char == asterisk {
			pos--
			continue
		}
		pos++
		sByte[pos] = byte(char)

	}

	fmt.Println(string(sByte[:pos+1]))

	return string(sByte[:pos+1])
}
