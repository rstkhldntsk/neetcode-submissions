func rotate(nums []int, k int) {
	helper := func(l, r int) {
		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	k %= len(nums)
	helper(0, len(nums)-1-k)
	helper(len(nums)-k, len(nums)-1)
	helper(0, len(nums)-1)
}
