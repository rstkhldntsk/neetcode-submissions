func groupAnagrams(strs []string) [][]string {
	hashes := make(map[[26]int][]string)
	for i := range strs {
		hash := [26]int{}
		word := strs[i]
		for j := range word {
			hash[word[j]-'a']++
		}
		if _, ok := hashes[hash]; ok {
			hashes[hash] = append(hashes[hash], word)
		} else {
			hashes[hash] = []string{word}
		}
	}
	result := [][]string{}
	for _, group := range hashes {
		result = append(result, group)
	}
	return result
}
