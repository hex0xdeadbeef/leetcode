package easy

func isValidParentheses(s string) bool {

	stack := []byte{}

	for i := range s {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
		default:
			if len(stack) != 0 && stack[len(stack)-1] == toFind(s[i]) {
				stack = stack[:len(stack)-1]
				continue
			}

			return false
		}
	}

	return len(stack) == 0

}

func toFind(b byte) byte {
	switch b {
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	default:
		return ' '
	}
}
