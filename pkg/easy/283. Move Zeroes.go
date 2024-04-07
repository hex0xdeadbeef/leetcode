package easy

// https://leetcode.com/problems/move-zeroes/?envType=study-plan-v2&envId=leetcode-75

func moveZeroes(nums []int) {

	var (
		result   = nums[:0]
		curIndex = -1
	)

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			result = append(result, nums[i])
			curIndex++
		}
	}

	copy(nums[curIndex+1:], make([]int, len(nums)-(curIndex+1)))

}
