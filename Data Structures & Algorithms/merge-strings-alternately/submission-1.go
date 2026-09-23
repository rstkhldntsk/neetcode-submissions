func mergeAlternately(word1 string, word2 string) string {
	i, k := 0, 0
	result := make([]byte, len(word1)+len(word2))
	for i < len(word1) || i < len(word2) {
		if i < len(word1) {
			result[k] = word1[i]
			k++
		}
		if i < len(word2) {
			result[k] = word2[i]
			k++
		}
		i++
	}
	return string(result)
}
