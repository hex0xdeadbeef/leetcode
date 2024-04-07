package middle

import "slices"

// https://leetcode.com/problems/determine-if-two-strings-are-close/?envType=study-plan-v2&envId=leetcode-75

func CloseStrings(word1 string, word2 string) bool {

	if len(word1) != len(word2) {
		return false
	}

	var (
		letters1, letters2 = make(map[rune]int, len(word1)), make(map[rune]int, len(word1))
		counts1, counts2   []int
	)

	for _, c := range word1 {
		letters1[c]++
	}

	for _, c := range word2 {
		letters2[c]++
	}

	if len(letters1) != len(letters2) {
		return false
	}

	for k, _ := range letters1 {
		if _, ok := letters2[k]; !ok {
			return false
		}
	}

	counts1, counts2 = make([]int, 0, len(letters1)), make([]int, 0, len(letters2))
	for k, n := range letters1 {
		counts1, counts2 = append(counts1, n), append(counts2, letters2[k])
	}

	slices.Sort(counts1)
	slices.Sort(counts2)

	for i, _ := range counts1 {
		if counts1[i] != counts2[i] {
			return false
		}
	}

	return true

}

/*
"cabbba"
"abbccc"
*/
