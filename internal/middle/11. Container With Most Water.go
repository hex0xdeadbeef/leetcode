package middle

// https://leetcode.com/problems/container-with-most-water/description/?envType=study-plan-v2&envId=leetcode-75

func maxArea(height []int) int {
	var (
		l, r = 0, len(height) - 1

		curArea, resultingArea int
	)

	for r-l >= 1 {
		curArea = min(height[l], height[r]) * (r - l)

		resultingArea = max(resultingArea, curArea)

		switch {
		case height[l] > height[r]:
			r--

		case height[r] > height[l]:
			l++
		}

	}

	return resultingArea

}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}
