package easy

// https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k/description/

func CountKDifference(nums []int, k int) int {
	var (
		occurencies = make(map[int]int, len(nums))
		cnt         int
	)

	for _, v := range nums {
		occurencies[v]++
	}

	for _, v := range nums {
		if n, ok := occurencies[v-k]; ok {
			cnt += n
		}
	}

	return cnt
}
