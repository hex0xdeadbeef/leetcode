package easy

// https://leetcode.com/problems/fair-candy-swap/submissions/1245541846/

// type empty struct{}

func FairCandySwap(aliceSizes []int, bobSizes []int) []int {
	const (
		divider = 2
	)

	var (
		aliceUnique, bobUnique = make(map[int]empty, len(aliceSizes)), make(map[int]empty, len(bobSizes))
		aliceSum, bobSum       = fillUnique(aliceSizes, aliceUnique), fillUnique(bobSizes, bobUnique)

		desirableNum = (bobSum + aliceSum) / divider
	)

	for k, _ := range aliceUnique {
		aliceDiff := aliceSum - k
		if _, ok := bobUnique[desirableNum-aliceDiff]; ok {
			return []int{k, desirableNum - aliceDiff}
		}

	}

	return nil

}

func fillUnique(nums []int, unique map[int]empty) int {

	var (
		sum int
	)
	for _, num := range nums {
		unique[num] = empty{}

		sum += num
	}

	return sum
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
