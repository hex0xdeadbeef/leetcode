package middle

// https://leetcode.com/problems/generate-parentheses/description/

func generateParenthesis(n int) []string {
	var (
		res = make([]string, 0, 2*n)

		rec func(openParenthesesCnt, closedParenthesesCnt int, curStr string)
	)

	rec = func(openParenthesesCnt, closedParenthesesCnt int, curStr string) {
		if len(curStr) == 2*n {
			res = append(res, curStr)
			return
		}

		if openParenthesesCnt < n {
			rec(openParenthesesCnt+1, closedParenthesesCnt, curStr+"(")
		}

		if closedParenthesesCnt < openParenthesesCnt {
			rec(openParenthesesCnt, closedParenthesesCnt+1, curStr+")")
		}
	}

	rec(0, 0, "")

	return res
}
