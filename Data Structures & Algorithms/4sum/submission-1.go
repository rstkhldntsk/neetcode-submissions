func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, len(nums)-1
			for l < r {
				curSum := nums[i] + nums[j] + nums[l] + nums[r]
				if curSum == target {
					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
                    for l < r && nums[l] == nums[l-1] {
                        l++
                    }
				} else if curSum > target {
					r--
				} else {
					l++
				}
			}
		}
	}
	return result
}
