package easy

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	var h, n int

	for {
		for ; h < len(haystack) && haystack[h] != needle[0]; h++ {
		}

		if h == len(haystack) {
			return -1
		}

		hStart := h

		for ; h < len(haystack) && n < len(needle) && haystack[h] == needle[n]; h, n = h+1, n+1 {
		}

		if haystack[hStart:h] == needle {
			return hStart
		}

		h, n = hStart+1, 0
	}
}
