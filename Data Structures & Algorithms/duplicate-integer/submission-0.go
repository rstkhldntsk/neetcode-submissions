func hasDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for i := range nums {
		if seen[nums[i]] {
			return true
		}
		seen[nums[i]] = true
	}
	return false
}
