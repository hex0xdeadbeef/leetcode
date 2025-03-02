package easy

// https://leetcode.com/problems/unique-number-of-occurrences/description/?envType=study-plan-v2&envId=leetcode-75

func uniqueOccurrences(arr []int) bool {

	var (
		occurrences      = make(map[int]int, len(arr))
		occurrencesCount map[int]empty
	)

	for _, n := range arr {
		occurrences[n]++
	}

	occurrencesCount = make(map[int]empty, len(occurrences))

	for _, v := range occurrences {
		if _, ok := occurrencesCount[v]; ok {
			return false
		}

		occurrencesCount[v] = empty{}
	}

	return true
}
