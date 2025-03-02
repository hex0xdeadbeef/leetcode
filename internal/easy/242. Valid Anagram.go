package easy

// https://leetcode.com/problems/valid-anagram/description/

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	var (
		sBytesCnts = make(map[byte]int, len(s))
	)

	for i := 0; i < len(s); i++ {
		sBytesCnts[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		sBytesCnts[t[i]]--

		if v, ok := sBytesCnts[t[i]]; ok && v == 0 {
			delete(sBytesCnts, t[i])
		}
	}

	return len(sBytesCnts) == 0
}
