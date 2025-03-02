package middle

// https://leetcode.com/problems/k-th-smallest-prime-fraction/?envType=daily-question&envId=2024-05-10

import (
	"sort"
)

type PrimeFraction struct {
	numerator   int
	denominator int

	fraction float64
}

type PrimeFractions []PrimeFraction

func (a PrimeFractions) Len() int           { return len(a) }
func (a PrimeFractions) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PrimeFractions) Less(i, j int) bool { return a[i].fraction < a[j].fraction }

func kthSmallestPrimeFraction(arr []int, k int) []int {
	const (
		idxPad = 1
	)

	var (
		primeFractions = make([]PrimeFraction, 0, len(arr)*len(arr))
	)

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			primeFractions = append(primeFractions, PrimeFraction{numerator: arr[i], denominator: arr[j], fraction: float64(arr[i]) / float64(arr[j])})
		}
	}

	sort.Sort(PrimeFractions(primeFractions))

	return []int{primeFractions[k-idxPad].numerator, primeFractions[k-idxPad].denominator}
}
