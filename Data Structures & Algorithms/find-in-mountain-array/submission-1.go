/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
	l, r := 0, mountainArr.length()-1
	for l < r {
		m := l + (r-l)/2
		mid := mountainArr.get(m)
		next := mountainArr.get(m+1)
		if mid < next {
			l = m + 1
		} else {
			r = m
		}
	}
	peak := l

	l, r = 0, peak-1
	for l <= r {
		m := l + (r-l)/2
		mid := mountainArr.get(m)
		if mid < target {
			l = m + 1
		} else if mid > target {
			r = m - 1
		} else {
			return m
		}
	}

	l, r = peak, mountainArr.length()-1
	for l <= r {
		m := l + (r-l)/2
		mid := mountainArr.get(m)
		if mid > target {
			l = m + 1
		} else if mid < target {
			r = m - 1
		} else {
			return m
		}
	}

	return -1
}
