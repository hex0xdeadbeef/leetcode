package easy

func removeDuplicates(nums []int) int {
	var (
		numsCopy            = nums[:1]
		copiedIdx, elemsCnt int
	)

	for i := 1; i < len(nums); i++ {
		if numsCopy[copiedIdx] != nums[i] {
			numsCopy = append(numsCopy, nums[i])
			copiedIdx++
			elemsCnt++
		}
	}

	return len(numsCopy)
}
