package middle

func findRepeatedDnaSequences(s string) []string {
	const windowSize = 10
	var res []string

	if len(s) < windowSize {
		return res
	}

	m := make(map[string]int, len(s)/windowSize)

	for i := 0; i < len(s)-windowSize+1; i++ {
		m[s[i:i+windowSize]]++
	}

	for k, v := range m {
		if v < 2 {
			continue
		}
		res = append(res, k)
	}

	return res
}
