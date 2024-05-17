package middle

// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/submissions/1224263888/?envType=study-plan-v2&envId=leetcode-75

type empty struct{}

func maxVowels(s string, k int) int {

	var (
		vowels   = map[byte]empty{'a': {}, 'i': {}, 'o': {}, 'e': {}, 'u': {}}
		curCount int
		maxCount int
	)

	for i := 0; i < k; i++ {
		if _, ok := vowels[s[i]]; ok {
			maxCount++
		}
	}

	curCount = maxCount

	for i := 1; i+k-1 < len(s); i++ {
		if _, ok := vowels[s[i-1]]; ok {
			curCount--
		}

		if _, ok := vowels[s[i+k-1]]; ok {
			curCount++
		}

		if curCount > maxCount {
			maxCount = curCount
		}

	}

	return maxCount
}
