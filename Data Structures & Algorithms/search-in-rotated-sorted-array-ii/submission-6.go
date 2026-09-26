func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return true
		}
		if nums[l] < nums[m] { // left sorted
			if nums[l] <= target && target < nums[m] { // target in left
				r = m - 1
			} else {
				l = m + 1
			}
		} else if nums[l] > nums[m] { // right sorted
			if nums[m] < target && target <= nums[r] { // target in right
				l = m + 1
			} else {
				r = m -1
			}
		} else {
			l++
		}
	}
	return false
}
