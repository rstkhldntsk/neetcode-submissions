func trap(height []int) int {
	maxLeft := make([]int, len(height))
	maxL := height[0]
	for i := 1; i < len(height); i++ {
		maxLeft[i] = maxL
		maxL = max(maxL, height[i])
	}
	
	maxRight := make([]int, len(height))
	maxR := height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		maxRight[i] = maxR
		maxR = max(maxR, height[i])
	}
	
	maxTrap := 0
	for i := 0; i < len(maxLeft); i++ {
		maxTrap += max(min(maxLeft[i], maxRight[i])-height[i], 0)
	}
	return maxTrap
}
