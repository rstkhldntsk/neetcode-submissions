func groupAnagrams(strs []string) [][]string {
	hashes := make(map[[26]int][]string)
	for _, word := range strs {
		hash := [26]int{}
		for _, c := range word {
			hash[c-'a']++
		}
		hashes[hash] = append(hashes[hash], word)
	}

	grouped := [][]string{}
	for _, group := range hashes {
		grouped = append(grouped, group)
	}
	return grouped
}
