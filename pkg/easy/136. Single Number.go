package easy

// https://leetcode.com/problems/single-number/

func singleNumber(nums []int) int {

	type empty struct{}
	seen := make(map[int]int, len(nums)/2)

	for _, v := range nums {
		seen[v]++
	}

	for k, v := range seen {
		if v == 1 {
			return k
		}
	}

	panic("UNREACHABLE")
}
