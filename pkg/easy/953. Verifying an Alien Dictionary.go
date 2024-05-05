package easy

// https://leetcode.com/problems/verifying-an-alien-dictionary/

func isAlienSorted(words []string, order string) bool {

	const (
		alphabet = "abcdefghijklmnopqrstuvwxyz"
	)

	var (
		oldToNewAlphabet = make(map[byte]byte, len(alphabet))
	)

	for i := 0; i < len(alphabet); i++ {
		oldToNewAlphabet[order[i]] = alphabet[i]
	}

	var (
		usualWord []byte
	)

	for i, word := range words {
		usualWord = make([]byte, len(word))

		for i := 0; i < len(word); i++ {
			usualWord[i] = oldToNewAlphabet[word[i]]
		}

		words[i] = string(usualWord)
	}

	for i := 0; i < len(words)-1; i++ {
		if words[i] > words[i+1] {
			return false
		}
	}

	return true

}
