package middle

// https://leetcode.com/problems/evaluate-reverse-polish-notation/description/

import "strconv"

func evalRPN(tokens []string) int {
	const (
		plus, minus, multiply, divide = "+", "-", "*", "/"
	)

	var (
		opsOccurencies = map[string]empty{plus: {}, minus: {}, multiply: {}, divide: {}}
		valsStack      = make([]string, 0, len(tokens)/2)

		performOp = func(a, b, op string) string {
			var (
				aNum, _ = strconv.Atoi(a)
				BNum, _ = strconv.Atoi(b)
			)

			switch op {
			case plus:
				return strconv.Itoa(aNum + BNum)
			case minus:
				return strconv.Itoa(aNum - BNum)
			case multiply:
				return strconv.Itoa(aNum * BNum)
			case divide:
				return strconv.Itoa(aNum / BNum)
			}

			panic("UNREACHABLE")
		}
	)

	for i := 0; i < len(tokens); i++ {
		_, ok := opsOccurencies[tokens[i]]
		if !ok {
			valsStack = append(valsStack, tokens[i])
			continue
		}

		a, b := valsStack[len(valsStack)-2], valsStack[len(valsStack)-1]
		valsStack = valsStack[:len(valsStack)-2]

		valsStack = append(valsStack, performOp(a, b, tokens[i]))
	}

	res, _ := strconv.Atoi(valsStack[len(valsStack)-1])

	return res
}
