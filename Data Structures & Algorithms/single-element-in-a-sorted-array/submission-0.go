func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	for r-l > 1 {
		m := l + (r-l)/2
		if nums[m] == nums[m+1] {
			if (m - l) % 2 == 1 {
				r = m - 1
			} else {
				l = m + 2
			}
		} else if nums[m] == nums[m-1] {
			if (r - m) % 2 == 1 {
				l = m + 1
			} else {
				r = m - 2
			}
		} else {
			return nums[m]
		}
	}
	return nums[l]
}