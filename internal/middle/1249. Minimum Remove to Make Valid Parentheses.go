package middle

// https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/description/

func minRemoveToMakeValid(s string) string {
	var (
		validIndices = make(map[int]struct{}, 1<<7)
		l, r         int
	)

Out:
	for {
		for ; l < len(s) && s[l] != '('; l++ {
		}
		if l == len(s) {
			break
		}

		if r < l {
			r = l + 1
		}

		for ; r < len(s) && s[r] != ')'; r++ {
		}

		switch r {
		case len(s):
			break Out

		default:
			validIndices[l], validIndices[r] = struct{}{}, struct{}{}
		}
		l++
		r++

		if r > len(s) {
			break
		}

	}

	var (
		res = make([]byte, 0, 1<<7)
	)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', ')':
			if _, ok := validIndices[i]; ok {
				res = append(res, s[i])
			}
		default:
			res = append(res, s[i])
		}

	}
	return string(res)
}
