package easy

// https://leetcode.com/problems/maximum-length-substring-with-two-occurrences/description/

func maximumLengthSubstring(s string) int {

	var (
		window = make(map[byte]int, len(s))

		i, startIdx int

		maxLen int
	)

	for ; i < len(s); i++ {
		if v, ok := window[s[i]]; ok && v+1 > 2 {
			if i-startIdx > maxLen {
				maxLen = i - startIdx
			}

			for window[s[i]] == 2 {
				window[s[startIdx]]--
				startIdx++
			}

			window[s[i]]++

			continue
		}

		window[s[i]]++

	}

	if i-startIdx > maxLen {
		maxLen = i - startIdx
	}

	return maxLen
}
