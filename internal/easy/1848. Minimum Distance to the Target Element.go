package easy

// https://leetcode.com/problems/minimum-distance-to-the-target-element/submissions/1500740460/

func getMinDistance(nums []int, target int, start int) int {
	var l, r = -1, -1

	for i := start; i < len(nums); i++ {
		if nums[i] == target {
			r = i
			break
		}
	}

	for i := start; i >= 0; i-- {
		if nums[i] == target {
			l = i
			break
		}
	}

	if l == -1 {
		return abs(r - start)
	}

	if r == -1 {
		return abs(l - start)
	}

	return min(abs(r-start), abs(l-start))

}
