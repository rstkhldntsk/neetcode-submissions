func twoSum(nums []int, target int) []int {
    hash := make(map[int]int)
	for i := range nums {
		if idx, ok := hash[target-nums[i]]; ok {
			return []int{idx, i}
		}
		hash[nums[i]] = i
	}
	return []int{}
}
