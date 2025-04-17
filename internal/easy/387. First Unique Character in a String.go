package easy

func firstUniqChar(s string) int {
	arr := [26]int{}

	for _, b := range []byte(s) {
		arr[b-'a']++
	}

	for i, b := range []byte(s) {
		if arr[b-'a'] == 1 {
			return i
		}
	}

	return -1
}
