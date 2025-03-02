package easy

// https://leetcode.com/problems/find-greatest-common-divisor-of-array/	

func findGCD(nums []int) int {

	const (
		firstElemIdx = 0
	)

	var (
		minVal, maxVal int = nums[firstElemIdx], nums[firstElemIdx]
	)

	for _, v := range nums {
		if v < minVal {
			minVal = v
			continue
		}

		if v > maxVal {
			maxVal = v
			continue
		}
	}

	return gcd(maxVal, minVal)

}

// func gcd(a, b int) int {
// 	if b == 0 {
// 		return a
// 	}

// 	return gcd(b, a%b)
// }
