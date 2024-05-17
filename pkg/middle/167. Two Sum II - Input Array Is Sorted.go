package middle

func twoSum(numbers []int, target int) []int {
	var (
		l, r   int = 0, len(numbers) - 1
		oldDif int
	)

	for numbers[l]+numbers[r] != target {
		oldDif = (numbers[l] + numbers[r]) - target

		if oldDif < 0 {
			l++
			continue
		}
		r--

	}

	return []int{l + 1, r + 1}
}
