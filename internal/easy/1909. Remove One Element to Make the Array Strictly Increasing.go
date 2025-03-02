package easy

// https://leetcode.com/problems/remove-one-element-to-make-the-array-strictly-increasing/description/

func canBeIncreasing(nums []int) bool {
	var (
		lCount, rCount int

		prev int
	)

	prev = nums[0]
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] <= nums[i] {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] <= prev {
					lCount++
				}
			}
		}

		prev = nums[i+1]
	}

	prev = nums[len(nums)-1]
	for i := len(nums) - 1; i > 1; i-- {
		if nums[i-1] >= nums[i] {
			for j := i - 1; j >= 0; j-- {
				if nums[j] >= prev {
					rCount++
				}
			}
		}

		prev = nums[i-1]
	}

	if lCount <= 1 || rCount <= 1 {
		return true
	}

	return false
}
