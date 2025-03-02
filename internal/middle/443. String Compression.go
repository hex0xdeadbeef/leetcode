package middle

import (
	"strconv"
)

// https://leetcode.com/problems/string-compression/description/?envType=study-plan-v2&envId=leetcode-75

func compress(chars []byte) int {

	const (
		sizeExtensionCoefficient = 2
		initialOffset            = 1
	)

	var (
		result = chars[:0]

		curChar      = chars[0]
		curCharCount int
	)

	for i := 0; i < len(chars); i++ {
		if curChar != chars[i] {
			insert(curChar, curCharCount, &result)

			curChar, curCharCount = chars[i], 0
		}

		curCharCount++
	}

	insert(curChar, curCharCount, &result)

	return len(result)

}

func insert(char byte, count int, sl *[]byte) {
	*sl = append(*sl, char)

	if count != 1 {
		*sl = append(*sl, []byte(strconv.Itoa(count))...)
	}

}
