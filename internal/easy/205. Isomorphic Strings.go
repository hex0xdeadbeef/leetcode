package easy

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	fromStoT := make(map[byte]byte)

	var (
		sP, tP int
	)

	for ; sP < len(s); sP, tP = sP+1, tP+1 {
		b, ok := fromStoT[s[sP]]
		if !ok {
			if isValPresent(t[tP], fromStoT) {
				return false
			}

			fromStoT[s[sP]] = t[tP]
			continue
		}

		if t[tP] != b {
			return false
		}
	}

	return true
}

func isValPresent(b byte, m map[byte]byte) bool {
	for _, v := range m {
		if v == b {
			return true
		}
	}

	return false
}
