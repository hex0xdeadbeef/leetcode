package middle

import . "leetcode/pkg/datastructures/doublylinkedlist"

func characterReplacement(s string, k int) int {
	var (
		usedLts = getUsedLetters(s)

		maxSequenceLen int
	)

	if k == 0 {
		return findMaxLetterSubsequenceLen(s)
	}

	for rep := range usedLts {
		maxSequenceLen = max(maxSequenceLen, findMaxSubstitutedSequenceLength(rep, k, s))
	}

	return maxSequenceLen
}

func findMaxSubstitutedSequenceLength(rep byte, k int, str string) int {
	var (
		repIdx, repCnt     = make([]int, 0, 1<<6), 0
		l, r           int = 0, 0

		maxLen int
	)

	for r < len(str) {

		if str[r] != rep {
			switch {
			case repCnt < k:
				repIdx = append(repIdx, r)
				repCnt++

			case repCnt == k:
				if curLen := r - l; curLen > maxLen {
					maxLen = curLen
				}
				l, repIdx = repIdx[0]+1, repIdx[1:]

				repIdx = append(repIdx, r)

			}

		}

		r++
	}

	if curLen := r - l; curLen > maxLen {
		maxLen = curLen
	}

	return maxLen
}

func findMaxLetterSubsequenceLen(str string) int {
	var (
		maxLen = 1

		l, r = 0, 1
	)

	for r < len(str) {

		for r < len(str) && str[r] == str[l] {
			r++
		}

		if curLen := r - l; curLen > maxLen {
			maxLen = curLen
		}

		l, r = r, r+1

	}

	if curLen := r - l; curLen > maxLen {
		maxLen = curLen
	}

	return maxLen
}

func getUsedLetters(s string) map[byte]struct{} {
	var (
		m = make(map[byte]struct{}, 1<<5)
	)

	for i := range s {
		m[s[i]] = struct{}{}
	}

	return m
}

func characterReplacementRepeat(s string, k int) int {
	var (
		alphabet       [26]int
		maxReplacement int
	)

	for i := 0; i < len(s); i++ {
		alphabet[s[i]-'A']++
	}

	for i, v := range alphabet {
		if v == 0 {
			continue
		}

		maxReplacement = max(maxReplacement, countWithLetter(s, 'A'+byte(i), k))
	}

	return maxReplacement
}

func countWithLetter(s string, letter byte, k int) int {
	var l, r, maxRange int
	list := New()

	for ; r < len(s) && (s[r] == letter || k > 0); r++ {
		if s[r] != letter && k > 0 {
			list.InsertTail(r)
			k--
		}
	}
	maxRange = r - l

	for ; r < len(s); r++ {
		if s[r] == letter {
			continue
		}

		if curRange := r - l - 1; curRange > maxRange {
			maxRange = curRange
		}

		if !list.IsEmpty() {
			l = list.PopHead().Val
			list.InsertTail(r)
			continue
		}

		l = r
	}

	if curRange := r - l - 1; curRange > maxRange {
		maxRange = curRange
	}

	return maxRange
}
