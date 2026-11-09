func hasDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	for i := range nums {
		seen[nums[i]] = struct{}{}
	}
	return len(seen) != len(nums)
}
