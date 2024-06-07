package easy

// https://leetcode.com/problems/score-of-a-string/description/?envType=daily-question&envId=2024-06-01

func scoreOfString(s string) int {

	var (
		score int
	)

	for i := 0; i < len(s)-1; i++ {
		score += abs(int(s[i]) - int(s[i+1]))
	}

	return score

}
