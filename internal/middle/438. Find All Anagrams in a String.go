package middle

func findAnagrams(s string, p string) []int {
	var res []int

	if len(p) > len(s) {
		return nil
	}

	mS, mP := make(map[byte]int, len(p)), make(map[byte]int, len(s))
	for _, b := range []byte(p) {
		mP[b]++
	}

	for i := 0; i < len(p); i++ {
		mS[s[i]]++
	}

	var i = len(p)
	for ; i < len(s); i++ {
		if mapsEqualFunc(mP, mS, func(v1, v2 int) bool { return v1 == v2 }) {
			res = append(res, i-len(p))
		}

		mS[s[i-len(p)]]--
		mS[s[i]]++
	}

	if mapsEqualFunc(mP, mS, func(v1, v2 int) bool { return v1 == v2 }) {
		res = append(res, i-len(p))
	}

	return res
}

func mapsEqualFunc(m1, m2 map[byte]int, f func(v1, v2 int) bool) bool {
	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; !ok || !f(v1, v2) {
			return false
		}
	}

	return true
}

func findAnagramsDiff(s string, p string) []int {
	var (
		res []int

		arrS, arrP [26]int
		diff       int
	)

	if len(s) < len(p) {
		return res
	}

	for i := 0; i < len(p); i++ {
		arrS[s[i]-'a']++
		arrP[p[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		if arrS[i] != arrP[i] {
			diff++
		}
	}

	for i := len(p); i < len(s); i++ {
		if diff == 0 {
			res = append(res, i-len(p))
		}

		oldCh := s[i-len(p)] - 'a'
		arrS[oldCh]--
		switch {
		case arrS[oldCh] == arrP[oldCh]:
			diff--

		case arrS[oldCh]+1 == arrP[oldCh]:
			diff++
		}

		newCh := s[i] - 'a'
		arrS[newCh]++
		switch {
		case arrS[newCh] == arrP[newCh]:
			diff--

		case arrS[newCh]-1 == arrP[newCh]:
			diff++
		}
	}

	if diff == 0 {
		res = append(res, len(s)-len(p))
	}

	return res
}

func findAnagramsNew(s string, p string) []int {
	if len(p) > len(s) {
		return nil
	}

	pArr := [26]int{}
	for _, b := range []byte(p) {
		pArr[b-'a']++
	}

	sArr := [26]int{}
	for i := 0; i < len(p); i++ {
		sArr[s[i]-'a']++
	}

	var res []int
	i := len(p)
	for ; i < len(s); i++ {
		if pArr == sArr {
			res = append(res, i-len(p))
		}

		sArr[s[i-len(p)]-'a']--
		sArr[s[i]-'a']++
	}

	if pArr == sArr {
		res = append(res, i-len(p))
	}

	return res
}
