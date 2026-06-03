func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	l, r := 0, len(height)-1
	highestL, highestR := height[l], height[r]
	maxTrap := 0
	for l < r {
		if highestL < highestR {
			l++
			maxTrap += max(highestL-height[l], 0)
			highestL = max(highestL, height[l])
		} else {
			r--
			highestR = max(highestR, height[r])
			maxTrap += highestR - height[r]
		}
	}
	return maxTrap
}
