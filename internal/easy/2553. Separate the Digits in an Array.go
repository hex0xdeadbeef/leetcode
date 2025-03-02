package easy

// https://leetcode.com/problems/separate-the-digits-in-an-array/description/

import (
	"strconv"
)

func separateDigits(nums []int) []int {

	const (
		lenMultiplier = 5
	)

	var (
		separatedDigits = make([]int, 0, len(nums)*lenMultiplier)

		curDigits = make([]int, lenMultiplier)

		splitDigits = func(num int) int {
			const (
				pad = 1
			)

			var (
				strNum = strconv.Itoa(num)
				count  int
			)

			for i, _ := range strNum {
				curDigits[i], _ = strconv.Atoi(strNum[i : i+pad])
				count++
			}

			return count
		}
	)

	for _, v := range nums {
		for i := 0; i < splitDigits(v); i++ {
			separatedDigits = append(separatedDigits, curDigits[i])
		}
	}

	return separatedDigits
}
