package easy

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/

func maxProfit(prices []int) int {
	const (
		oldLowIdx = iota
		oldTopIdx
		newLowIdx
		newTopIdx

		placeHolder = 1e9
	)

	var (
		pricesPool = []int{oldLowIdx: prices[0], oldTopIdx: -1, newTopIdx: placeHolder}
	)
	prices = append(prices, prices[len(prices)-1])

	// Initialize oldTop price
	for i := 0; i < len(prices) && i <= oldTopIdx; i++ {
		pricesPool[oldTopIdx] = prices[i]
	}
	pricesPool[newLowIdx], pricesPool[newTopIdx] = pricesPool[oldLowIdx], pricesPool[oldTopIdx]

	for i := oldTopIdx; i < len(prices); i++ {
		if (pricesPool[newTopIdx] - pricesPool[newLowIdx]) > (pricesPool[oldTopIdx] - pricesPool[oldLowIdx]) {
			pricesPool[oldTopIdx], pricesPool[oldLowIdx] = pricesPool[newTopIdx], pricesPool[newLowIdx]
		}

		if prices[i] < pricesPool[newLowIdx] {
			pricesPool[newLowIdx], pricesPool[newTopIdx] = prices[i], prices[i+1]
			continue
		}

		if prices[i] > pricesPool[newTopIdx] {
			pricesPool[newTopIdx] = prices[i]
			continue
		}

	}

	if (pricesPool[newTopIdx] - pricesPool[newLowIdx]) > (pricesPool[oldTopIdx] - pricesPool[oldLowIdx]) {
		pricesPool[oldTopIdx], pricesPool[oldLowIdx] = pricesPool[newTopIdx], pricesPool[newLowIdx]
	}

	if pricesPool[oldTopIdx]-pricesPool[oldLowIdx] < 0 {
		return 0
	}

	return pricesPool[oldTopIdx] - pricesPool[oldLowIdx]
}
