func rob(nums []int) int {
	a, b, c := 0, 0, nums[0]
	for i := 1; i < len(nums); i++ {
		a, b, c = b, c, nums[i]+max(a, b)
	}
	return max(b, c)
}
