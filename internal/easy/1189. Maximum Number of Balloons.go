package easy

// https://leetcode.com/problems/maximum-number-of-balloons/description/

func maxNumberOfBalloons(text string) int {
	var (
		subtractions = map[byte]int{
			'b': 1,
			'a': 1,
			'l': 2,
			'o': 2,
			'n': 1,
		}

		symbolsCounts map[byte]int = make(map[byte]int, len(text))

		count int
	)

	for i := 0; i < len(text); i++ {
		if _, ok := subtractions[text[i]]; ok {
			symbolsCounts[text[i]]++
		}
	}

out:
	for {

		for k, v1 := range subtractions {
			if v2, _ := symbolsCounts[k]; v2-v1 < 0 {
				break out
			}

			symbolsCounts[k] -= v1
		}

		count++
	}

	return count
}
