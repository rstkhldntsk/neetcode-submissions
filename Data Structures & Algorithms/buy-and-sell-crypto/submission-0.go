func maxProfit(prices []int) int {
    mprofit := 0
    for l, r := 0, 1; r < len(prices); r++ {
        profit := prices[r]-prices[l]
        mprofit = max(mprofit, profit)
        if profit < 0 {
            l = r
        }
    }
    return mprofit
}
