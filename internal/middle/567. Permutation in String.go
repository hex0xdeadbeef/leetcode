package middle

// https://leetcode.com/problems/permutation-in-string/description/

func checkInclusion(word string, str string) bool {
	var (
		window          = make(map[byte]int, 1<<5)
		wordLetterToCnt = make(map[byte]int, len(str))
	)

	if len(word) > len(str) {
		return false
	}

	for i := range word {
		wordLetterToCnt[word[i]]++
	}

	for i := 0; i < len(word); i++ {
		window[str[i]]++
	}

	for i := 0; i < len(str); i++ {
		if isMatched(wordLetterToCnt, window) {
			return true
		}

		window[str[i]]--
		if v, _ := window[str[i]]; v == 0 {
			delete(window, str[i])
		}

		if len(word)+i == len(str) {
			return false
		}
		window[str[len(word)+i]]++
	}

	return false
}

func isMatched(super, sub map[byte]int) bool {
	if len(super) != len(sub) {
		return false
	}

	for k, v1 := range super {
		v2, ok := sub[k]
		if !ok {
			return false
		}

		if v1 != v2 {
			return false
		}
	}

	return true
}

func checkInclusionSmart(s1, s2 string) bool {
	var (
		diff       int
		arr1, arr2 [26]int
	)

	for _, b := range s1 {
		arr1[b-'a']++
	}

	for i := 0; i < len(s1); i++ {
		arr2[s2[i]-'a']++
	}

	for i, v := range arr1 {
		if v != arr2[i] {
			diff++
		}
	}

	for i := len(s1); i < len(s2); i++ {
		if diff == 0 {
			return true
		}
		
		oldCh := s2[i-len(s1)] - 'a'
		arr2[oldCh]--
		if arr1[oldCh] == arr2[oldCh] {
			diff--
		} else if arr2[oldCh]+1 == arr1[oldCh] {
			diff++
		}

		newCh := s2[i] - 'a'
		arr2[newCh]++
		if arr1[newCh] == arr2[newCh] {
			diff--
		} else if arr2[newCh]-1 == arr1[newCh] {
			diff++
		}
	}

	return diff == 0
}
