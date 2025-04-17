package middle

func totalFruit(fruits []int) int {
	if len(fruits) <= 2 {
		return len(fruits)
	}

	var (
		lastTypeIdx, l, r, maxCount int
		swapped                     bool
	)

	t1, t2 := fruits[0], fruits[0]

	for ; r < len(fruits); r++ {
		if t1 == t2 && fruits[r] != t1 {
			t2, lastTypeIdx = fruits[r], r
			swapped = !swapped
			continue
		}

		if swapped && (t1 == fruits[r] || t2 == fruits[r]) && fruits[r] != fruits[lastTypeIdx] {
			lastTypeIdx = r
		}

		if t1 != t2 && fruits[r] != t1 && fruits[r] != t2 {
			break
		}

	}

	if curCount := r - l; curCount > maxCount {
		maxCount = curCount
	}

	for ; r < len(fruits); r++ {
		if t1 != t2 && fruits[r] != t1 && fruits[r] != t2 {
			if curCount := r - l; curCount > maxCount {
				maxCount = curCount
			}

			l, lastTypeIdx, t1, t2 = lastTypeIdx, r, fruits[lastTypeIdx], fruits[r]
			continue
		}

		if fruits[r] != fruits[lastTypeIdx] {
			lastTypeIdx = r
		}

	}

	if curCount := r - l; curCount > maxCount {
		maxCount = curCount
	}

	return maxCount
}
