func minEatingSpeed(piles []int, h int) int {
	canEat := func(k int) bool {
		hours := 0
		for _, pile := range piles {
			hours += pile / k
			if pile%k != 0 {
				hours++
			}
		}
		return hours <= h
	}
	
	maxSpeed := 0
	for _, pile := range piles {
		maxSpeed = max(maxSpeed, pile)
	}
	
	l, r := 0, maxSpeed
	for r-l > 1 {
		k := l + (r-l)/2
		if canEat(k) {
			r = k
		} else {
			l = k
		}
	}
	return r
}
