package easy

import "fmt"

// https://leetcode.com/problems/can-place-flowers/?envType=study-plan-v2&envId=leetcode-75

func canPlaceFlowers(flowerbed []int, n int) bool {

	var (
		l = 0
		r int

		dif         int
		subtraction int
	)

	if len(flowerbed) == 1 {
		if flowerbed[0] == 0 && n != 0 {
			n--
		}
	}

	for i := l; i < len(flowerbed) && n > 0; i++ {
		if flowerbed[i] == 0 {
			l = i
			r = getRightBoundary(l, flowerbed)
			dif = r - l

			switch {
			case l == 0 && r == len(flowerbed):
				if r%2 == 1 {
					subtraction = (r/2 + 1)
				} else {
					subtraction = r / 2
				}

			case l == 0 || r == len(flowerbed):
				subtraction = dif / 2
			default:
				switch {
				case dif%2 == 0:
					subtraction = (dif/2 - 1)
				default:
					subtraction = dif / 2
				}
			}

			n -= subtraction
			l = r
			i = l
		}
	}

	fmt.Println(n == 0)

	return n <= 0

}

func getRightBoundary(l int, segment []int) int {
	r := l

	for i := l; i < len(segment); i++ {
		if segment[i] == 0 {
			r++
			continue
		}

		break
	}

	return r
}
