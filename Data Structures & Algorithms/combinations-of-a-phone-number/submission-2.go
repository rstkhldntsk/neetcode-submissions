func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	alphabet := map[byte][]string {
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	combinations := make([]string, 0)
	var backtrack func(prefix string, idx int)
	backtrack = func(prefix string, idx int) {
		if len(prefix) == len(digits) {
			combinations = append(combinations, prefix)
			return
		}

		letters := alphabet[digits[idx]]
		for i := range letters {
			backtrack(prefix+letters[i], idx+1)
		}
	}
	backtrack("", 0)

	return combinations
}
