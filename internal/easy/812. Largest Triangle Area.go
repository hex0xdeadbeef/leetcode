package easy

import "math"

// https://leetcode.com/problems/largest-triangle-area/description/

func largestTriangleArea(points [][]int) float64 {

	var (
		maxArea float64

		d1, d2, d3      float64
		halvedPerimeter float64
	)

	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			for k := 0; k < len(points); k++ {
				d1, d2, d3 = getDistance(points[i], points[j]), getDistance(points[j], points[k]), getDistance(points[i], points[k])
				halvedPerimeter = getHalvedPerimeter(d1, d2, d3)

				if curArea := getArea(d1, d2, d3, halvedPerimeter); curArea > maxArea {
					maxArea = curArea
				}
			}
		}
	}

	return maxArea
}

func getDistance(p1, p2 []int) float64 {
	return math.Sqrt(math.Pow(float64(p1[0]-p2[0]), 2) + math.Pow(float64(p1[1]-p2[1]), 2))
}

func getHalvedPerimeter(d1, d2, d3 float64) float64 {
	return (d1 + d2 + d3) / 2
}

func getArea(d1, d2, d3 float64, halvedPerimeter float64) float64 {
	return math.Sqrt(halvedPerimeter *
		(halvedPerimeter - d1) *
		(halvedPerimeter - d2) *
		(halvedPerimeter - d3))
}
