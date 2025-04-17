package easy

import "strings"

func longestCommonPrefix(strs []string) string {
	shortest := strs[0]
	for i := range strs {
		if len(strs[i]) <= len(shortest) {
			shortest = strs[i]
		}
	}

	for i := len(shortest); i > 0; i-- {
		allInclude := true

		for j := range strs {
			if strings.HasPrefix(strs[j], shortest[:i]) {
				continue
			}

			allInclude = false
		}

		if allInclude {
			return shortest[:i]
		}

	}

	return ""

}
