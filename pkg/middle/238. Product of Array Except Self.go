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

func productExceptSelfEasy(nums []int) []int {
	var (
		res = make([]int, len(nums))

		curProduct int = 1
	)

	for i, v := range nums {
		res[i] = curProduct
		curProduct *= v
	}

	curProduct = 1
	for i := len(nums) - 1; i > -1; i-- {
		res[i] = curProduct * res[i]
		curProduct *= nums[i]
	}

	return res
}
