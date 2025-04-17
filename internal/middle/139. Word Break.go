package middle

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]struct{})
	for _, word := range wordDict {
		wordSet[word] = struct{}{}
	}

	memo := make(map[string]bool)
	return canBreak(s, wordSet, memo)
}

func canBreak(s string, wordSet map[string]struct{}, memo map[string]bool) bool {
	if s == "" {
		return true
	}

	if val, exists := memo[s]; exists {
		return val
	}

	for i := 1; i <= len(s); i++ {
		if _, ok := wordSet[s[:i]]; ok {
			if canBreak(s[i:], wordSet, memo) {
				memo[s] = true
				return true
			}
		}
	}

	memo[s] = false
	return false
}
