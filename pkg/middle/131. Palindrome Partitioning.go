package middle

// https://leetcode.com/problems/palindrome-partitioning/description/?envType=daily-question&envId=2024-05-22

func partition(s string) [][]string {
	var (
		length = len(s)
		res    = make([][]string, 0, length)

		parts = make([]string, 0)

		strPartitioning func(idx int)
	)

	strPartitioning = func(idx int) {
		if idx == length {
			newParts := make([]string, len(parts))
			copy(newParts, parts)
			res = append(res, newParts)
			return
		}

		for i := idx + 1; i <= length; i++ {
			if isPalindrome(s[idx:i]) {
				parts = append(parts, s[idx:i])
				strPartitioning(i)
				parts = parts[:len(parts)-1]
			}

		}
	}

	strPartitioning(0)

	return res

}

func twoPow(pow int) int {
	return 1 << pow
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
