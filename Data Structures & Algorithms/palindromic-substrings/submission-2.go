func countSubstrings(s string) int {
	countPalindromes := func(l, r int) int {
		cnt := 0
		for l >= 0 && r < len(s) && s[l] == s[r] {
			cnt++
			l--
			r++
		}
		return cnt
	}
	res := 0
	for i := range s {
		res += countPalindromes(i, i)
		res += countPalindromes(i, i+1)
	}
	return res
}