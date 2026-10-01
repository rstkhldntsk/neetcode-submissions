func lengthOfLongestSubstring(s string) int {
	hash := [128]bool{}
	maxLen := 0
	for l, r := 0, 0; r < len(s); r++ {
		for hash[s[r]] {
			hash[s[l]] = false
			l++
		}
		hash[s[r]] = true
		maxLen = max(maxLen, r-l+1)
	}
	return maxLen
}
