func maxArea(heights []int) int {
	result := 0
	l, r := 0, len(heights)-1
	for l < r {
		result = max(result, min(heights[l], heights[r])*(r-l))
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}
	return result
}
