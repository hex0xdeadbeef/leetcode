package easy

// https://leetcode.com/problems/find-lucky-integer-in-an-array/description/

func findLucky(arr []int) int {
	var (
		frequencies map[int]int = make(map[int]int, len(arr))
	)

	for _, v := range arr {
		frequencies[v]++
	}

	var (
		maxLucky int = -1
	)
	for k, v := range frequencies {
		if k == v && v > maxLucky {
			maxLucky = v
		}
	}

	if maxLucky != -1 {
		return maxLucky
	}

	return -1
}
