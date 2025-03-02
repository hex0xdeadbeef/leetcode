package middle

func MaxSatisfied(customers []int, grumpy []int, minutes int) int {
	var res, maxS, curS int

	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			res += customers[i]
		}
	}

	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 {
			curS += customers[i]
		}
	}
	maxS = curS

	for i := minutes; i < len(customers); i++ {
		if grumpy[i] == 1 {
			curS += customers[i]
		}

		if grumpy[i-minutes] == 1 {
			curS -= customers[i-minutes]
		}

		if curS > maxS {
			maxS = curS
		}
	}

	return res + maxS
}
