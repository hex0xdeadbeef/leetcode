package easy

func isHappy(n int) bool {
	m := map[int]struct{}{}

	var newNum int
	for {
		newNum = 0

		for n != 0 {
			newNum += (n % 10) * (n % 10)
			n /= 10
		}

		_, ok := m[newNum]
		if ok {
			return false
		}

		m[newNum] = struct{}{}

		if newNum == 1 {
			return true
		}

		n = newNum
	}
}
