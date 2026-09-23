func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		} else if nums[l] <= nums[m] { // leftSorted
			if nums[l] <= target && target < nums[m] { // targetInLeftHalf
				r = m - 1
			} else { // targetInRightHalf
				l = m + 1
			}
		} else { // rightSorted
			if nums[m] < target && target <= nums[r] {  // targetInRightHalf
				l = m + 1
			} else { // targetInLeftHalf
				r = m - 1
			}
		}
	}
	return -1
}
