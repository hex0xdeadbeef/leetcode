	package middle

	// https://leetcode.com/problems/maximize-distance-to-closest-person/

	const (
		pad = 1

		emptyS, takenS = 0, 1
	)

	func maxDistToClosest(seats []int) int {

		const (
			firstElementIdx, secondElementIdx = 0, 1

			two, oneReminder = 2, 1
		)

		var (
			maxDistance int
		)

		if seats[0] == emptyS {
			var (
				leftDistance int
			)

			leftDistance, seats = checkLeft(seats)
			maxDistance = max(maxDistance, leftDistance)
		}

		if seats[len(seats)-pad] == emptyS {
			var (
				rightDistance int
			)
			rightDistance, seats = checkRight(seats)
			maxDistance = max(maxDistance, rightDistance)
		}

		var (
			r int = secondElementIdx

			curEmptySeatsCount int
		)

		for r < len(seats) {

			if seats[r] == takenS {
				for r < len(seats) && seats[r] == takenS {
					r++
				}
			}

			for r < len(seats) && seats[r] == emptyS {
				curEmptySeatsCount++
				r++
			}

			switch curEmptySeatsCount % two {
			case oneReminder:
				maxDistance = max(maxDistance, curEmptySeatsCount/two+pad)

			default:
				maxDistance = max(maxDistance, curEmptySeatsCount/two)
			}

			curEmptySeatsCount = 0

		}

		return maxDistance
	}

	func checkLeft(seats []int) (int, []int) {
		var (
			idx   int
			count int = -1
		)

		for ; idx < len(seats); idx++ {
			count++

			if seats[idx] == takenS {
				break
			}
		}

		return count, seats[idx:]
	}

	func checkRight(seats []int) (int, []int) {
		var (
			idx   int
			count int = -1
		)

		for idx = len(seats) - 1; idx >= 0; idx-- {
			count++

			if seats[idx] == takenS {
				break
			}
		}

		return count, seats[:idx+1]
	}

	// func max(a, b int) int {
	// 	if a > b {
	// 		return a

	// 	}
	// 	return b
	// }
