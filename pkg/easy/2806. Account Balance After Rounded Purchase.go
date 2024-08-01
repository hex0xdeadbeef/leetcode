package easy

// https://leetcode.com/problems/account-balance-after-rounded-purchase/description/

func accountBalanceAfterPurchase(purchaseAmount int) int {
	var (
		roundedAmount = purchaseAmount - (purchaseAmount % 10)
	)

	if purchaseAmount == 0 {
		return 100
	}

	if remainder := purchaseAmount % 10; remainder == 0 || remainder < 5 {
		return 100 - roundedAmount
	}

	return 100 - (roundedAmount + 10)
}
