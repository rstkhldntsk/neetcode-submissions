func maxProduct(nums []int) int {
	result := nums[0]
	prefix := 1
	suffix := 1
	for i := range nums {
		prefix *= nums[i]
		suffix *= nums[len(nums)-i-1]
		result = max(result, max(prefix, suffix))
		if prefix == 0 {
			prefix = 1
		}
		if suffix == 0 {
			suffix = 1
		}
	}
	return result
}
