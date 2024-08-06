package easy

// https://leetcode.com/problems/kth-distinct-string-in-an-array/description/?envType=daily-question&envId=2024-08-05

func kthDistinct(arr []string, k int) string {
	var (
		indices = make(map[string]int, len(arr))
		cnts    = make([]int, len(arr))
	)

	for i, s := range arr {
		if _, ok := indices[s]; ok {
			continue
		}

		indices[s] = i
	}

	for _, v := range arr {
		cnts[indices[v]]++
	}

	for i, v := range cnts {
		if v == 1 {
			k--
		}

		if k == 0 {
			return arr[i]
		}
	}

	return ""
}
