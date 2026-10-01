func containsNearbyDuplicate(nums []int, k int) bool {
	hash := make(map[int]int)
	
	for i, num := range nums {
		if j, ok := hash[num]; ok && i-j <= k {
			return true
		}
		hash[num] = i
	}

	return false
}
