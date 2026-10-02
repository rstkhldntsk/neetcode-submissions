func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-1
	for r-l >= k {
		lDist := abs(arr[l] - x)
		rDist := abs(arr[r] - x)
		if lDist > rDist {
			l++
		} else {
			r--
		}
	}
	return arr[l : r+1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
