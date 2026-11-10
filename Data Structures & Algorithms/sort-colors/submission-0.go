func sortColors(nums []int) {
	lt := 0
	gt := len(nums)-1
	i := lt
	for i <= gt {
		if nums[i] == 0 {
			nums[i], nums[lt] = nums[lt], nums[i]
			i++
			lt++
		} else if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			nums[i], nums[gt] = nums[gt], nums[i]
			gt--
		}
	}
}
