func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	half := total / 2
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	l, r := 0, len(nums1)-1
	for {
		i := l + (r-l)/2
		j := half - i
		
		aLeft := math.MinInt
		if i > 0 {
			aLeft = nums1[i-1]
		}
		aRight := math.MaxInt
		if i < len(nums1) {
			aRight = nums1[i]
		}
		bLeft := math.MinInt
		if j > 0 {
			bLeft = nums2[j-1]
		}
		bRight := math.MaxInt
		if j < len(nums2) {
			bRight = nums2[j]
		}
		
		// correct partition
		if aLeft <= bRight && bLeft <= aRight {
			// odd
			if total%2 == 1 {
				return float64(min(aRight, bRight))
			}
			return float64(max(aLeft, bLeft)+min(aRight, bRight)) / 2.
		} else if aLeft > bRight {
			r = i - 1
		} else {
			l = i + 1
		}
	}
}
