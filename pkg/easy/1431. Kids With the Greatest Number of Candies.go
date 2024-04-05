package easy

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {

	var (
		max = candies[0]

		indicators = make([]bool, len(candies))
	)

	for _, n := range candies {
		if n > max {
			max = n
		}
	}

	for i, n := range candies {
		if n+extraCandies >= max {
			indicators[i] = true
		}
	}

	fmt.Println(indicators)

	return indicators
}
