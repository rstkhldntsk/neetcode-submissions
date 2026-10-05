func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-1
	for l < r {
		m := l + (r-l)/2
		if arr[m] < x {
			l = m+1
		} else {
			r = m
		}
	}

	l = l-1
	r = l+1
	for r-l-1 < k {
		if l < 0 {
			r++
		} else if r >= len(arr) {
			l--
		} else if abs(arr[l]-x) <= abs(arr[r]-x) {
			l--
		} else {
			r++
		}
	}

	return arr[l+1:r]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
