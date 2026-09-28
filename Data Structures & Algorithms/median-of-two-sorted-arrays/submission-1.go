func findMedianSortedArrays(a []int, b []int) float64 {
	if len(a) > len(b) {
		return findMedianSortedArrays(b, a)
	}
	total := len(a) + len(b)
	half := total/2
	l, r := 0, len(a)-1
	for {
		m := l + (r-l)/2
		m2 := half - m

		aLeft := math.MinInt
		if m-1 >= 0 {
			aLeft = a[m-1]
		}
		aRight := math.MaxInt
		if m < len(a) {
			aRight = a[m]
		}
		bLeft := math.MinInt
		if m2-1 >= 0 {
			bLeft = b[m2-1]
		}
		bRight := math.MaxInt
		if m2 < len(b) {
			bRight = b[m2]
		}

		if aLeft <= bRight && bLeft <= aRight { // corect partition
			if total % 2 == 1 {
				return float64(min(aRight, bRight))
			}
			return float64(max(aLeft, bLeft) + min(aRight, bRight)) / 2.
		} else if bLeft > aRight {
			l = m + 1
		} else {
			r = m - 1
		}
	}
}
