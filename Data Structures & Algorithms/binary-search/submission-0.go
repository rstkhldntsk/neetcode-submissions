func search(nums []int, target int) int {
	l, r := -1, len(nums)
	for r-l > 1 {
		m := l + (r-l)/2
		if nums[m] > target {
			r = m
		} else if nums[m] < target {
			l = m
		} else {
			return m
		}
	}
	return -1
}
