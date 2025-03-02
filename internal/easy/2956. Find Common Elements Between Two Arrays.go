package easy

// type empty struct{}

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	var (
		unique1, unique2 = make(map[int]empty, len(nums1)), make(map[int]empty, len(nums2))
		count1, count2   int
	)

	for _, v := range nums1 {
		unique1[v] = empty{}
	}

	for _, v := range nums2 {
		unique2[v] = empty{}

		if _, ok := unique1[v]; ok {
			count2++
		}

	}

	for _, v := range nums1 {
		if _, ok := unique2[v]; ok {
			count1++
		}
	}

	return []int{count1, count2}
}
