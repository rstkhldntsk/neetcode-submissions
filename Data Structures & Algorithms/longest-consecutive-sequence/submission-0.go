func longestConsecutive(nums []int) int {
	res := 0
	mp := make(map[int]int)
	for _, num := range nums {
		if mp[num] == 0 {
			mp[num] = mp[num-1] + mp[num+1] + 1
			mp[num-mp[num-1]] = mp[num]
			mp[num+mp[num+1]] = mp[num]
			res = max(res, mp[num])
		}
	}
	return res
}
