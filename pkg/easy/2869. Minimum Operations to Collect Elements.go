package easy

// https://leetcode.com/problems/minimum-operations-to-collect-elements/description/

func minOperations(nums []int, k int) int {
	var (
		m = make(map[int]struct{}, k)

		opsCnt int
	)

	for i := 1; i <= k; i++ {
		m[i] = struct{}{}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if len(m) == 0 {
			break
		}

		delete(m, nums[i])
		opsCnt++
	}

	return opsCnt
}
