func trap(height []int) int {
	maxLeft := make([]int, len(height))
	maxL := height[0]
	for i := 1; i < len(height); i++ {
		maxLeft[i] = maxL
		maxL = max(maxL, height[i])
	}
	
	maxRight := 0
	maxTrap := 0
	for i := len(height) - 1; i >= 0; i-- {
		maxTrap += max(min(maxRight, maxLeft[i])-height[i], 0)
		maxRight = max(maxRight, height[i])
	}
	return maxTrap
}
