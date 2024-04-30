package easy

// https://leetcode.com/problems/distance-between-bus-stops/description/

func DistanceBetweenBusStops(distance []int, start int, destination int) int {
	const (
		pad = 1
	)

	var (
		n = len(distance)

		clockwise = func() int {
			var (
				sum     int
				curStop = start
			)

			for curStop != destination {
				sum += distance[curStop]
				curStop = (curStop + pad) % n
			}

			return sum

		}

		counterclockwise = func() int {
			var (
				sum     int
				curStop = start - pad
			)

			for {
				if curStop < 0 {
					curStop = n - pad
				}
				sum += distance[curStop]

				if curStop == destination {
					break
				}

				curStop--
			}

			return sum

		}
	)

	return min(clockwise(), counterclockwise())

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
