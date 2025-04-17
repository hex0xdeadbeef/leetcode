package middle

// https://leetcode.com/problems/group-anagrams/description/

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {

	const (
		mapSize   = 512
		sliceSize = 4
	)

	var (
		sortedASCUniqueWords = make(map[string][]string, mapSize)

		letters []string

		result [][]string
	)

	for _, v := range strs {

		letters = strings.Split(v, "")
		sort.Strings(letters)
		sortedWord := strings.Join(letters, "")

		if _, ok := sortedASCUniqueWords[sortedWord]; !ok {
			sortedASCUniqueWords[sortedWord] = make([]string, 0, sliceSize)

			sortedASCUniqueWords[sortedWord] = append(sortedASCUniqueWords[sortedWord], v)
			continue
		} else {
			sortedASCUniqueWords[sortedWord] = append(sortedASCUniqueWords[sortedWord], v)
		}
	}

	result = make([][]string, 0, len(sortedASCUniqueWords))
	for _, v := range sortedASCUniqueWords {
		result = append(result, v)
	}

	return result
}

func groupAnagramsNew(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for i := range strs {
		arr := [26]int{}

		for _, b := range []byte(strs[i]) {
			arr[b-'a']++
		}

		m[arr] = append(m[arr], strs[i])
	}

	res := [][]string{}
	for _, s := range m {
		res = append(res, s)
	}

	return res
}
