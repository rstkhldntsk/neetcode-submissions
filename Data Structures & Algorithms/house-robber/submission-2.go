func rob(nums []int) int {
	prev, cur := 0, 0
	for i := range nums {
		prev, cur = cur, max(cur, nums[i]+prev)
	}
	return cur
}
