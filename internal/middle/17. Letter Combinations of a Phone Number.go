package middle

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	digitToLetters := map[byte][]byte{}

	var (
		iterationsCount int
		letter          byte = 'a'
	)

	for i := 2; i <= 9; i++ {
		iterationsCount = 3
		if i == 7 || i == 9 {
			iterationsCount++
		}

		for range iterationsCount {
			digitToLetters['0'+byte(i)] = append(digitToLetters['0'+byte(i)], letter)
			letter++
		}
	}

	res := &[]string{}

	letterCombinationsTraverse(digits, 0, nil, digitToLetters, res)

	return *res
}

func letterCombinationsTraverse(digits string, curIdx int, curS []byte, digitToLetters map[byte][]byte, res *[]string) {
	if curIdx == len(digits) {
		*res = append(*res, string(curS))
		return
	}

	for _, b := range digitToLetters[digits[curIdx]] {
		curS = append(curS, b)
		letterCombinationsTraverse(digits, curIdx+1, curS, digitToLetters, res)
		curS = curS[:len(curS)-1]
	}

}
