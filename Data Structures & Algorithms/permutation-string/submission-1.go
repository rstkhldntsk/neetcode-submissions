func checkInclusion(s1 string, s2 string) bool {
	hash := [26]int{}
	for i := range s1 {
		hash[s1[i]-'a']++
	}

	for l, r := 0, 0; r < len(s2); r++ {
		hash[s2[r]-'a']--
		if r-l+1 > len(s1) {
			hash[s2[l]-'a']++
			l++
		}
		if hash == [26]int{} {
			return true
		}
	}
	return false
}
