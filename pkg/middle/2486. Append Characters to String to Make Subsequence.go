package middle

// https://leetcode.com/problems/append-characters-to-string-to-make-subsequence/description/?envType=daily-question&envId=2024-06-03

func appendCharacters(s string, t string) int {
	var (
		sP, tP int
	)

	for ; tP < len(t) && sP < len(s); sP++ {
		if s[sP] == t[tP] {
			tP++
		}

	}

	return len(t) - tP
}
