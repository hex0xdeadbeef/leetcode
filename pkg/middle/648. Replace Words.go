package middle

// https://leetcode.com/problems/replace-words/description/?envType=daily-question&envId=2024-06-07

import (
	"sort"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	var (
		words        = strings.Fields(sentence)
		dictionaries = genDictionaries(dictionary)
		replacement  string
	)

	for i, w := range words {
		firstLetter := []rune(w)[0]

		_, ok := dictionaries[firstLetter]

		if !ok {
			replacement = w
		}

		replacement = findLongestPrefix(w, dictionaries[firstLetter])

		words[i] = replacement
	}

	return strings.Join(words, " ")
}

func genDictionaries(dict []string) map[rune][]string {
	var (
		dictionaries = make(map[rune][]string)
	)

	for _, word := range dict {
		firstLetter := []rune(word)[0]

		_, ok := dictionaries[firstLetter]

		if !ok {
			dictionaries[firstLetter] = make([]string, 0, 16)
		}

		dictionaries[firstLetter] = append(dictionaries[firstLetter], word)
	}

	for _, v := range dictionaries {
		sort.Strings(v)
	}

	return dictionaries
}

func findLongestPrefix(word string, prefixesByLetter []string) string {
	for _, p := range prefixesByLetter {
		if strings.HasPrefix(word, p) {
			return p
		}
	}

	return word
}
