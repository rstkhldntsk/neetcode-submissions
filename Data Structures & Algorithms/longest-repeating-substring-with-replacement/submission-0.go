func characterReplacement(s string, k int) int {
	hash := [26]int{}
	res, freq := 0, 0
	for l, r := 0, 0; r < len(s); r++ {
		hash[s[r]-'A']++
		freq = max(freq, hash[s[r]-'A'])
		for r-l-freq >= k {
			hash[s[l]-'A']--
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}
