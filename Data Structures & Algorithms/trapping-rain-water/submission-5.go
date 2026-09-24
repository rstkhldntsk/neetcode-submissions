func trap(height []int) int {
	maxL, maxR := height[0], height[len(height)-1]
	l, r := 0, len(height)-1
	maxTrap := 0
	for l < r {
		if maxL < maxR {
			l++
			maxTrap += max(maxL-height[l], 0)
			maxL = max(maxL, height[l])
		} else {
			r--
			maxTrap += max(maxR-height[r], 0)
			maxR = max(maxR, height[r])
		}
	}
	return maxTrap
}
