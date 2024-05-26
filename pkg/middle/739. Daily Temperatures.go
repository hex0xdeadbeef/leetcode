package middle

// https://leetcode.com/problems/daily-temperatures/description/

func dailyTemperatures(temperatures []int) []int {
	type (
		t struct {
			val int
			idx int
		}
	)

	var (
		stack, res = make([]t, 0, len(temperatures)), make([]int, len(temperatures))
	)

	for i := len(temperatures) - 1; i > -1; i-- {
		switch {
		case len(stack) == 0:
			res[i] = 0

		case temperatures[i] >= stack[len(stack)-1].val:
			for len(stack) != 0 && temperatures[i] >= stack[len(stack)-1].val {
				stack = stack[:len(stack)-1]
			}

			if len(stack) == 0 {
				res[i] = 0
			} else {
				res[i] = stack[len(stack)-1].idx - i
			}

		default:
			res[i] = stack[len(stack)-1].idx - i

		}

		stack = append(stack, t{val: temperatures[i], idx: i})

	}

	return res

}
