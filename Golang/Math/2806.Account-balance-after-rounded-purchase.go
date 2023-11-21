package main

func accountBalanceAfterPurchase(purchaseAmount int) int {
	digit := purchaseAmount % 10
	if digit < 5 {
		purchaseAmount = purchaseAmount / 10 * 10
	} else {
		purchaseAmount = (purchaseAmount/10 + 1) * 10
	}
	return 100 - purchaseAmount
}
