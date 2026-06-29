func longestPalindrome(s string) string {
	resL, resLen := 0, 0
	for i := 0; i < len(s); i++ {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > resLen {
				resL = l
				resLen = r - l + 1
			}
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > resLen {
				resL = l
				resLen = r - l + 1
			}
			l--
			r++
		}
	}
	return s[resL : resL+resLen]
}
