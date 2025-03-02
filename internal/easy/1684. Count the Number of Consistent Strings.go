package easy

// https://leetcode.com/problems/count-the-number-of-consistent-strings/

// type empty struct{}

func сountConsistentStrings(allowed string, words []string) int {

	var (
		allowedSymbols = fillTheSetOfSymbolsAllowed(allowed)

		count int
	)

outer:
	for _, word := range words {

		for i := 0; i < len(word); i++ {
			if _, ok := allowedSymbols[word[i]]; !ok {
				continue outer
			}

		}

		count++
	}

	return count
}

func fillTheSetOfSymbolsAllowed(allowed string) map[byte]empty {

	var (
		result = make(map[byte]empty, len(allowed))
	)

	for i := 0; i < len(allowed); i++ {
		result[allowed[i]] = empty{}
	}

	return result
}
