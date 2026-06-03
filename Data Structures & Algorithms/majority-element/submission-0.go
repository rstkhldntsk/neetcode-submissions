func majorityElement(nums []int) int {
    cand := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == cand {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			cand = nums[i]
			cnt = 1
		}
	}
	return cand
}
