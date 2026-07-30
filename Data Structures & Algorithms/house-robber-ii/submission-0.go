func rob(nums []int) int {
    var rob1 func (houses []int) int 
    rob1 = func (houses []int) int {
        prev, cur := 0, 0
        for i := range houses {
            prev, cur = cur, max(cur, houses[i]+ prev)
        }
        return cur
    }
    return max(nums[0], max(rob1(nums[:len(nums)-1]), rob1(nums[1:])))
}