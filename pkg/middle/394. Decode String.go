package middle

// https://leetcode.com/problems/decode-string/?envType=study-plan-v2&envId=leetcode-75

import (
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {

	const (
		open  = '['
		close = ']'
	)

	var (
		strStack, numericStack = make([]string, 0, 1024), make([]int, 0, 1024)

		tempStr   string
		numberStr string

		result string
	)

	strStack = append(strStack, "")
	for _, r := range s {
		if unicode.IsLetter(r) {
			tempStr += string(r)
			continue
		}

		if unicode.IsNumber(r) {
			numberStr += string(r)
			continue
		}

		if r == open {
			switch num, err := strconv.Atoi(numberStr); {
			case err == nil:
				numericStack = append(numericStack, num)
			default:
				numericStack = append(numericStack, 1)
			}
			numberStr = ""

			if len(tempStr) != 0 {
				strStack[len(strStack)-1] += tempStr
			}
			strStack = append(strStack, "")
			tempStr = ""
		}

		if r == close {
			num := numericStack[len(numericStack)-1]
			numericStack = numericStack[:len(numericStack)-1]

			strStack[len(strStack)-1] += tempStr
			tempStr = ""

			if len(strStack) != 1 {
				strStack[len(strStack)-2] += strings.Repeat(strStack[len(strStack)-1], num)
				strStack = strStack[:len(strStack)-1]
				continue
			}

			strStack[len(strStack)-1] = strings.Repeat(strStack[len(strStack)-1], num)
		}

	}

	for _, s := range strStack {
		result += s

	}
	result += tempStr

	return result

}
