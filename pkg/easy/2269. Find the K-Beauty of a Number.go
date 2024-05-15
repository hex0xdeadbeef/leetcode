package easy

// https://leetcode.com/problems/find-the-k-beauty-of-a-number/description/

import "strconv"

func DivisorSubstrings(num int, k int) int {
	var (
		strNum = strconv.Itoa(num)

		window string
		curNum int

		cnt int
	)

	for i := 0; i < k; i++ {
		window += string([]byte{strNum[i]})
	}

	curNum, _ = strconv.Atoi(window)
	if curNum != 0 && num%curNum == 0 {
		cnt++
	}

	for i := k; i < len(strNum); i++ {
		window = window[1:]
		window += string([]byte{strNum[i]})

		curNum, _ = strconv.Atoi(window)
		if curNum != 0 && num%curNum == 0 {
			cnt++
		}
	}

	return cnt

}
