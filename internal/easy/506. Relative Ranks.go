package easy

// https://leetcode.com/problems/relative-ranks/description/?envType=daily-question&envId=2024-05-08

import (
	"sort"
	"strconv"
)

type scorePair struct {
	idx int
	val int
}

type scorePairs []scorePair

func (a scorePairs) Len() int           { return len(a) }
func (a scorePairs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a scorePairs) Less(i, j int) bool { return a[i].val > a[j].val }

func findRelativeRanks(score []int) []string {

	const (
		idxPad = 1
	)

	var (
		ranks = []string{"Gold Medal", "Silver Medal", "Bronze Medal"}

		pairs, result = make([]scorePair, len(score)), make([]string, len(score))
	)

	for i, v := range score {
		pairs[i] = scorePair{idx: i, val: v}
	}
	sort.Sort(scorePairs(pairs))

	var ranksP int
	for i := 0; i < len(pairs) && i < len(ranks); i++ {
		result[pairs[i].idx] = ranks[ranksP]
		ranksP++
	}

	ranksP++

	for j := ranksP - idxPad; j < len(pairs); j++ {
		result[pairs[j].idx] = strconv.Itoa(ranksP)
		ranksP++
	}

	return result
}
