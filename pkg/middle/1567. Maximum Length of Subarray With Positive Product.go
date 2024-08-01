package middle

// https://leetcode.com/problems/maximum-length-of-subarray-with-positive-product/

type NumsGroup struct {
	base             []int
	prefixL, prefixR []int
}

func getMaxLen(nums []int) int {
	var (
		maxPositiveProductLength int
	)

	for i, v := range nums {
		switch {
		case v < 0:
			nums[i] = -1
		case v > 0:
			nums[i] = 1
		}
	}
	groups := getGroups(nums)

	calculatePrefixProducts(groups)
	for _, g := range groups {
		for i := len(g.prefixL) - 1; i >= 0; i-- {
			if g.prefixL[i] > 0 {
				if i > maxPositiveProductLength {
					maxPositiveProductLength = i
					break
				}
			}
		}

		for i := len(g.prefixR) - 1; i >= 0; i-- {
			if g.prefixR[i] > 0 {
				if i > maxPositiveProductLength {
					maxPositiveProductLength = i
					break
				}
			}
		}

	}

	return maxPositiveProductLength
}

func getGroups(nums []int) []NumsGroup {
	const (
		pad = 1

		prefixBase = 1
	)
	var (
		groups = make([]NumsGroup, 0, 1<<7)
		l, r   int
	)

	for l < len(nums) {

		for ; r < len(nums) && nums[r] != 0; r++ {
		}

		switch {
		case l != r:
			groups = append(groups, NumsGroup{base: nums[l:r], prefixL: make([]int, r-l+pad), prefixR: make([]int, r-l+pad)})
			groups[len(groups)-pad].prefixL[0], groups[len(groups)-pad].prefixR[0] = prefixBase, prefixBase
		}

		l, r = r+pad, r+pad
	}

	return groups
}

func calculatePrefixProducts(groups []NumsGroup) {
	const (
		pad = 1
	)

	for _, g := range groups {
		for i := 0; i < len(g.base); i++ {
			g.prefixL[i+1] = g.prefixL[i] * g.base[i]
		}

		var rPrefixIdx = 1
		for i := len(g.base) - 1; i >= 0; i-- {
			g.prefixR[rPrefixIdx] = g.prefixR[rPrefixIdx-pad] * g.base[i]
			rPrefixIdx++
		}
	}
}
