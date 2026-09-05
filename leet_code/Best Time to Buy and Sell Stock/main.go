package main

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for _, currentPrice := range prices {
		if currentPrice < minPrice {
			minPrice = currentPrice

		} else {
			curProfit := currentPrice - minPrice
			if curProfit > maxProfit {
				maxProfit = curProfit
			}
		}

	}
	return maxProfit

}
