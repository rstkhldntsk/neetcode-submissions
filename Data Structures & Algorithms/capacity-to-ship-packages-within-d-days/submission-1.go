func shipWithinDays(weights []int, days int) int {
	canShip := func(capacity int) bool {
		d := 0
		curLoad := 0
		for _, w := range weights {
			if w > capacity {
				return false
			}
			if curLoad == 0 || curLoad+w > capacity {
				d++
				curLoad = w
			} else {
				curLoad += w
			}
		}
		return d <= days
	}
	
	total := 0
	maxW := 0
	for _, w := range weights {
		total += w
		maxW = max(maxW, w)
	}
	
	l, r := maxW-1, total
	for r-l > 1 {
		capacity := l + (r-l)/2
		if canShip(capacity) {
			r = capacity
		} else {
			l = capacity
		}
	}
	return r
}
