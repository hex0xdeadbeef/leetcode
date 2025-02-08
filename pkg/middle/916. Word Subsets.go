package middle

func WordSubsets(words1 []string, words2 []string) []string {
	var (
		letterToNumber = make(map[byte]int, len(words2))

		res []string
	)

	for _, w := range words2 {
		for _, b := range []byte(w) {
			letterToNumber[b]++
		}
	}

	for _, word := range words1 {
		wordSet := getLettersSet(word)

		var balancer int

		for k, v := range letterToNumber {
			wordSet[k] -= v

			if wordSet[k] == 0 {
				balancer++
			}
		}

		if balancer == len(letterToNumber) {
			res = append(res, word)
		}
	}

	return res

}

func getLettersSet(s string) map[byte]int {
	m := make(map[byte]int, 26)

	for _, b := range []byte(s) {
		m[b]++
	}

	return m
}
