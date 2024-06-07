package easy

// https://leetcode.com/problems/longest-palindrome/description/?envType=daily-question&envId=2024-06-04

func longestPalindrome(s string) int {
	const (
		alphabetSize = 52
	)

	var (
		occToCnt = make(map[byte]int, alphabetSize)
		resL     int

		remainder, subtraction int
	)

	for i := 0; i < len(s); i++ {
		occToCnt[s[i]]++
	}

	for k, v := range occToCnt {
		remainder = v % 2
		switch remainder {
		case 1:
			subtraction = v - 1
		default:
			subtraction = v
		}
		resL += subtraction
		occToCnt[k] -= subtraction
	}

	for _, v := range occToCnt {
		if v == 1 {
			resL += 1
			break
		}
	}

	return resL
}
