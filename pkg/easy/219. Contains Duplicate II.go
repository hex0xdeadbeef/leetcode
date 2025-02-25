package easy

func containsNearbyDuplicate(nums []int, k int) bool {
	var l, r int

	m := make(map[int]int, k)

	for r = 0; r <= k && r < len(nums); r++ {
		m[nums[r]]++

		if v := m[nums[r]]; v == 2 {
			return true
		}
	}

	for ; r < len(nums); l, r = l+1, r+1 {
		m[nums[l]]--
		m[nums[r]]++

		if v := m[nums[r]]; v == 2 {
			return true
		}
	}

	return false
}
