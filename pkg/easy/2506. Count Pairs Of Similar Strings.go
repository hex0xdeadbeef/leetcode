package easy

// https://leetcode.com/problems/count-pairs-of-similar-strings/description/

func similarPairs(words []string) int {
	const (
		bitShiftPad = 'a'
	)

	var (
		pairsCount int

		wordsLetters = make([]uint32, len(words))
	)

	for i, word := range words {
		for j := 0; j < len(word); j++ {
			if curLetterBitShit := uint32(word[j] - bitShiftPad); (wordsLetters[i]>>curLetterBitShit)&1 == 0 {
				wordsLetters[i] += 1 << curLetterBitShit
			}
		}
	}

	for i := 0; i < len(wordsLetters); i++ {
		for _, v := range wordsLetters[i+1:] {
			if wordsLetters[i] == v {
				pairsCount++
			}
		}
	}

	return pairsCount
}
