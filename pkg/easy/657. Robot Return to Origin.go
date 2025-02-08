package easy

// https://leetcode.com/problems/robot-return-to-origin/

func judgeCircle(moves string) bool {
	var ud, lr int

	for _, r := range moves {
		switch r {
		case 'U':
			ud++
		case 'D':
			ud--
		case 'R':
			lr++
		case 'L':
			lr--
		}
	}

	return ud == 0 && lr == 0
}
