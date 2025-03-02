package easy

// https://leetcode.com/problems/contains-duplicate/

func containsDuplicate(nums []int) bool {

	var (
		seen = make(map[int]int, len(nums))
	)

	for _, v := range nums {
		seen[v]++

		if seen[v] == 2 {
			return true
		}
	}

	return false

}
