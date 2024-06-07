package easy

import "fmt"

// https://leetcode.com/problems/find-common-characters/?envType=daily-question&envId=2024-06-05

func commonChars(words []string) []string {
	const (
		pad          = 97
		alphabetSize = 26
	)
	var (
		matrix = make([][]int, len(words))

		res = make([]string, 0, alphabetSize)
	)

	for i, word := range words {
		matrix[i] = make([]int, alphabetSize)

		for _, r := range word {
			matrix[i][r-pad]++
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 0; j < alphabetSize; j++ {
			matrix[0][j] = min(matrix[0][j], matrix[i][j])
		}
	}

	for i := 0; i < alphabetSize; i++ {
		for j := 0; j < matrix[0][i]; j++ {
			res = append(res, fmt.Sprintf("%c", i+pad))
		}
	}

	return res

}
