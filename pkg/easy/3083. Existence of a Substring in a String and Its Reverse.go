package easy

// https://leetcode.com/problems/existence-of-a-substring-in-a-string-and-its-reverse/

// type empty struct{}

func isSubstringPresent(s string) bool {

	var (
		twoLengthSequences = make(map[string]empty, 32)
	)

	for i := 0; i < len(s)-1; i++ {

		if s[i] == s[i+1] {
			return true
		}

		twoLengthSequences[string([]byte{s[i], s[i+1]})] = empty{}
	}

	for i := 0; i < len(s)-1; i++ {
		if _, ok := twoLengthSequences[string([]byte{s[len(s)-1-i], s[len(s)-i-2]})]; ok {
			return true
		}
	}

	return false
}
