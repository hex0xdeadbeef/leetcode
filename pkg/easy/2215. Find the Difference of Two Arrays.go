package easy

// https://leetcode.com/problems/find-the-difference-of-two-arrays/?envType=study-plan-v2&envId=leetcode-75

func findDifference(nums1 []int, nums2 []int) [][]int {
	var (
		result     = make([][]int, 2)
		set1, set2 = make(map[int]empty, len(nums1)), make(map[int]empty, len(nums2))
	)

	for i := 0; i < len(nums1); i++ {
		set1[nums1[i]] = empty{}
	}

	for i := 0; i < len(nums2); i++ {
		set2[nums2[i]] = empty{}
	}

	result[0] = make([]int, 0, len(set1))
	result[0] = make([]int, 0, len(set2))

	for k := range set1 {

		if _, ok := set2[k]; !ok {
			result[0] = append(result[0], k)
		}

	}

	for k, _ := range set2 {

		if _, ok := set1[k]; !ok {
			result[1] = append(result[1], k)
		}

	}

	return result

}
