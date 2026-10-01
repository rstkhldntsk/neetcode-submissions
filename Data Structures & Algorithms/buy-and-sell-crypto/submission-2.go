func maxProfit(prices []int) int {
	maxP := 0
	for l, r := 0, 1; r < len(prices); r++ {
		if prices[r] > prices[l] {
			maxP = max(maxP, prices[r]-prices[l])
		} else {
			l = r
		}
	}
	return maxP
}
