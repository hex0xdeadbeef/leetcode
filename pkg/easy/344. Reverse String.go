package easy

// https://leetcode.com/problems/reverse-string/?envType=daily-question&envId=2024-06-02

func reverseString(s []byte) {
	var (
		l = len(s)
	)
	for i := 0; i < len(s)/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}
