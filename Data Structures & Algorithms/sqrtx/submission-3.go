func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	l, r := 0, x
	for r-l > 1 {
		m := l + (r-l)/2
		if m*m <= x {
			l = m
		} else {
			r = m
		}
	}
	return l
}
