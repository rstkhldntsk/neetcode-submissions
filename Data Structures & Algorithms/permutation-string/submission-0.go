func checkInclusion(s1 string, s2 string) bool {
	hash := [26]int{}
	for _, c := range s1 {
		hash[c-'a']++
	}
	l := 0
	for r, c := range s2 {
		hash[c-'a']--
		if hash == [26]int{} {
			return true
		}
		if r+1 >= len(s1) {
			hash[s2[l]-'a']++
			l++
		}
	}
	return false
}
