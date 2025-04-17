package middle

func NumberOfSubstrings(s string) int {

	type pair struct {
		subS string
		l, r int
	}

	var (
		l, r int
		arr  [3]int

		m = map[pair]struct{}{}
	)

	for ; r < len(s); r++ {
		if isABC(arr) {
			m[pair{subS: s[l:r], l: l, r: r}] = struct{}{}
		}

		arr[s[r]-'a']++
	}

	for r-l >= 1 {
		if isABC(arr) {
			m[pair{subS: s[l:r], l: l, r: r}] = struct{}{}
		}

		arr[s[l]-'a']--
		l++
	}

	l, r, arr = 0, 0, [3]int{}
	for ; r < len(s) && !isABC(arr); r++ {
		arr[s[r]-'a']++
	}

	// aaabca
	for r < len(s) {
		for ; l < len(s) && isABC(arr); l++ {
			arr[s[l]-'a']--
		}

		for ; r < len(s) && !isABC(arr); r++ {
			arr[s[r]-'a']++
		}

		m[pair{subS: s[l:r], l: l, r: r}] = struct{}{}

		for r := r; r < len(s); r++ {
			m[pair{subS: s[l:r], l: l, r: r}] = struct{}{}
		}
	}

	if isABC(arr) {
		m[pair{subS: s[l:r], l: l, r: r}] = struct{}{}
	}

	return len(m)
}

func isABC(arr [3]int) bool {
	for i := range arr {
		if arr[i] <= 0 {
			return false
		}
	}

	return true
}
