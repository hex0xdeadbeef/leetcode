package easy

func canConstruct(ransomNote string, magazine string) bool {
	arr := [26]int{}

	for _, b := range []byte(magazine) {
		arr[b-'a']++
	}

	for _, b := range []byte(ransomNote) {
		arr[b-'a']--

		if arr[b-'a'] < 0 {
			return false
		}
	}

	return true
}
