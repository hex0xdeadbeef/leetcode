package middle

// https://leetcode.com/problems/find-the-duplicate-number/description/

func findDuplicateNum(nums []int) int {
	var (
		slow, fast = 0, 0
	)

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	fast = 0
	for {
		fast = nums[fast]
		slow = nums[slow]

		if fast == slow {
			return slow
		}
	}
}
