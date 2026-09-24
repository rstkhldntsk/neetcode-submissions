func trap(height []int) int {
	maxL, maxR := height[0], height[len(height)-1]
	l, r := 0, len(height)-1
	maxTrap := 0
	for l < r {
		if maxL < maxR {
			l++
			maxL = max(maxL, height[l])
			maxTrap += maxL-height[l]
		} else {
			r--
			maxR = max(maxR, height[r])
			maxTrap += maxR-height[r]
		}
	}
	return maxTrap
}
