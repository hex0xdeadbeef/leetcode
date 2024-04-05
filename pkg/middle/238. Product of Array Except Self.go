package middle

// https://leetcode.com/problems/product-of-array-except-self/description/?envType=study-plan-v2&envId=leetcode-75

func productExceptSelf(nums []int) []int {

	var (
		l        = len(nums)
		prefixes = make([]int, l)
	)

	preMultiple := 1
	for i := 0; i < l; i++ {
		prefixes[i] = preMultiple
		preMultiple *= nums[i]
	}

	postMultiple := 1
	for i := l - 1; i >= 0; i-- {
		prefixes[i] *= postMultiple
		postMultiple *= nums[i]
	}

	return prefixes
}
