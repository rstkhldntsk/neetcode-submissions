func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash := [26]int{}
	for i := range s {
		hash[s[i]-'a']++
		hash[t[i]-'a']--
	}
	for i := range hash {
		if hash[i] != 0 {
			return false
		}
	}
	return true
}
