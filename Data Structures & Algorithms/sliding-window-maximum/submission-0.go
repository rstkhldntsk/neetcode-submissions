func maxSlidingWindow(nums []int, k int) []int {
    result := make([]int, 0, len(nums)-k+1)
	decreasingDeque := make([]int, 0)
	for r := 0; r < len(nums); r++ {
		for len(decreasingDeque) != 0 && nums[decreasingDeque[len(decreasingDeque)-1]] < nums[r] {
			decreasingDeque = decreasingDeque[:len(decreasingDeque)-1]
		}
		decreasingDeque = append(decreasingDeque, r)
		if decreasingDeque[0] == r-k {
			decreasingDeque = decreasingDeque[1:]
		}
		if r+1 >= k {
			result = append(result, nums[decreasingDeque[0]])
		}
	}
	return result
}
