func numSquares(n int) int {
    dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = n
	}
	
	for target := 1; target <= n; target++ {
		for i := 1; i*i <= target; i++ {
            sq := i*i
			dp[target] = min(dp[target], 1+dp[target-sq])
		}
	}
	return dp[n]
}
