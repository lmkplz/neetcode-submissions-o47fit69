func maxProfit(prices []int) int {
	res := 0
	profit := 0
	high := 0
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > high {
			high = prices[i]
		} else {
			profit = high - prices[i]

			if profit > res {
				res = profit
			}
		}
	}

	return res
}
