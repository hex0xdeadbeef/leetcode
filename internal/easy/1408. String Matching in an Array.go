package easy

import "strings"

func stringMatching(words []string) []string {

	var (
		m    = map[string]int{}
		resM = map[string]struct{}{}
	)

	for _, w := range words {
		m[w]++
	}

	for k, v := range m {

		switch {
		case v > 1:
			resM[k] = struct{}{}

		default:
			for kToCheck := range m {
				if k != kToCheck && strings.Contains(kToCheck, k) {
					resM[k] = struct{}{}
				}
			}
		}
	}

	var res []string

	for k := range resM {
		res = append(res, k)
	}

	return res
}
