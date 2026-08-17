func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	var backtrack func(prefix []string, idx int)
	backtrack = func(prefix []string, idx int) {
		if len(prefix) == 4 && idx == len(s) {
			result = append(result, strings.Join(prefix, "."))
			return
		}
		for i := idx; i < min(idx+3, len(s)); i++ {
			subs := s[idx : i+1]
			if len(subs) > 1 && subs[0] == '0' { // only one leading zero is allowed
				continue
			}
			
			n, _ := strconv.Atoi(subs)
			if n > 255 {
				continue
			}
			
			backtrack(append(prefix, subs), i+1)
		}
	}
	backtrack([]string{}, 0)
	return result
}
