package easy

// https://leetcode.com/problems/sum-of-squares-of-special-elements/submissions/1232118927/

func sumOfSquares(nums []int) int {
	var (
		r   = len(nums) / 2
		sum int
	)

	for i := 0; i < r; i++ {
		if len(nums)%(i+1) == 0 {
			sum += square(nums[i])
		}
	}

	sum += square(nums[len(nums)-1])

	return sum
}

func square(a int) int {
	return a * a
}
