package easy

import "math"

// https://leetcode.com/problems/construct-the-rectangle/description/

func constructRectangle(area int) []int {
	var (
		result      = []int{area, 1}
		minDiff int = 1e9
	)

	for i := area / 2; i >= 1; i-- {
		if area%i != 0 {
			continue
		}

		if i-(area/i) < 0 {
			break
		}

		if i-(area/i) < minDiff {
			result[0], result[1] = i, area/i
		}

	}

	return result
}

func constructRectangleFast(area int) []int {
	for m := int(math.Sqrt(float64(area))); ; m-- {
		if area%m == 0 {
			return []int{area / m, m}
		}
	}
}
