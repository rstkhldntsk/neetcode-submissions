func minSubArrayLen(target int, nums []int) int {
	result, curSum := math.MaxInt, 0
	for l, r := 0, 0; r < len(nums); r++ {
		curSum += nums[r]
		for curSum >= target {
			result = min(result, r-l+1)
			curSum -= nums[l]
			l++
		}
	}
	if result == math.MaxInt {
		return 0
	}
	return result
}
