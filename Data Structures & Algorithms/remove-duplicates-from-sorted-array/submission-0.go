func removeDuplicates(nums []int) int {
	l, r := 1, 1
	for ; r < len(nums); r++ {
		if nums[r] != nums[r-1] {
			nums[l] = nums[r]
			l++
		}
	}
	return l
}
