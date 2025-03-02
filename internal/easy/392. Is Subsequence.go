package easy

// https://leetcode.com/problems/is-subsequence/description/?envType=study-plan-v2&envId=leetcode-75

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s) > len(t) {
		return false
	}

	var (
		ps int
	)

	for _, c := range t {
		if ps == len(s) {
			break
		}

		if byte(c) == s[ps] {
			ps++
		}
	}

	return ps == len(s)
}
