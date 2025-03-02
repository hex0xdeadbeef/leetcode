package easy

// https://leetcode.com/problems/next-greater-element-i/description/

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	var (
		seen = make(map[int]int, len(nums2))
	)

	for i, n1 := range nums1 {

		if _, ok := seen[n1]; ok {
			nums1[i] = n1
		}
		ind := findIndex(n1, nums2)

		for ; ind < len(nums2); ind++ {
			if nums2[ind] > n1 {
				nums1[i] = nums2[ind]
				seen[ind] = nums2[ind]
				break

			}
		}

		if ind >= len(nums2) {
			nums1[i] = -1
			seen[n1] = -1
		}

	}

	return nums1
}

func findIndex(num int, values []int) int {
	for i, v := range values {
		if v == num {
			return i
		}
	}

	panic("UNREACHABLE")
}
