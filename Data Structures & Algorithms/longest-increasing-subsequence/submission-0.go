func lengthOfLIS(nums []int) int {
    dp := make([]int, len(nums))
    for i := range nums {
        curMax := 0
        for j := range i {
            if nums[i] > nums[j] && dp[j] > curMax {
                curMax = dp[j]
            }
        }
        dp[i] = curMax+1
    }
    res := 0
    for i := range dp {
        res = max(res, dp[i])
    }
    return res
}