package middle

// https://leetcode.com/problems/max-number-of-k-sum-pairs/description/?envType=study-plan-v2&envId=leetcode-75

func maxOperations(nums []int, k int) int {

	var (
		p           int
		filteredMap = make(map[int]int, len(nums))

		n1, n2 int
		c1, c2 int

		count int
	)

	// Filter and addition the numbers less than k
	for _, n := range nums {
		if n < k {
			nums[p] = n
			p++

			filteredMap[n]++
		}
	}

	for i := 0; i < p; i++ {
		n1, n2 = nums[i], k-nums[i]
		c1, c2 = filteredMap[n1], filteredMap[n2]

		if n1 == n2 {
			switch {
			case c1 >= 2:
				filteredMap[n1] -= 2
				count++
			default:
				delete(filteredMap, n1)
			}

		}

		if n1 != n2 {
			switch {
			case c1 > 0 && c2 > 0:
				filteredMap[n1]--
				filteredMap[n2]--
				count++
			default:
				delete(filteredMap, n1)
				delete(filteredMap, n2)
			}
		}

	}

	return count
}
