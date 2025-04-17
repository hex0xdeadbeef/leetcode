package middle

func longestPalindrome(s string) string {

	var (
		res    string
		maxLen int
	)

	for i := 0; i < len(s); i++ {
		for j := len(s); j >= 0; j-- {
			if dif := j - i; isPalindrome(s[i:j]) && dif > maxLen {
				res, maxLen = s[i:j], dif
			}
		}
	}

	return res

}
