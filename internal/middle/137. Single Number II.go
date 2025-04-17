package middle

func singleNumber(nums []int) int {
	m := map[int]int{}

	for i := range nums {
		m[nums[i]]++
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}

	return 0

}
